package gce

import (
	"context"

	"github.com/pharmer/cloud/pkg/apis/cloud/v1"
	"github.com/pharmer/cloud/pkg/util"
	"github.com/pharmer/pharmer/credential"
	"golang.org/x/oauth2/google"
	compute "google.golang.org/api/compute/v1"
)

type Client struct {
	Data           *GceData
	GceProjectID   string
	ComputeService *compute.Service
	Ctx            context.Context
}

type GceData v1.CloudProvider

func NewClient(gecProjectId, credentialFilePath string) (*Client, error) {
	g := &Client{
		GceProjectID: gecProjectId,
		Ctx:          context.Background(),
		Data:         &GceData{},
	}
	var err error
	g.ComputeService, err = getComputeService(g.Ctx, credentialFilePath)
	if err != nil {
		return nil, err
	}
	data, err := util.GetDataFormFile("gce")
	if err != nil {
		return nil, err
	}
	d := GceData(*data)
	g.Data = &d
	return g, nil
}

func (g *Client) GetName() string {
	return g.Data.Name
}

func (g *Client) GetCredentials() []v1.CredentialFormat {
	return g.Data.Spec.Credentials
}

func (g *Client) GetKubernetes() []v1.KubernetesVersion {
	return g.Data.Spec.Kubernetes
}

func (g *Client) GetRegions() ([]v1.Region, error) {
	req := g.ComputeService.Regions.List(g.GceProjectID)

	regions := []v1.Region{}
	err := req.Pages(g.Ctx, func(list *compute.RegionList) error {
		for _, region := range list.Items {
			res, err := ParseRegion(region)
			if err != nil {
				return err
			}
			regions = append(regions, *res)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return regions, err
}

func (g *Client) GetZones() ([]string, error) {
	req := g.ComputeService.Zones.List(g.GceProjectID)
	zones := []string{}
	err := req.Pages(g.Ctx, func(list *compute.ZoneList) error {
		for _, zone := range list.Items {
			zones = append(zones, zone.Name)
		}
		return nil
	})

	if err != nil {
		return nil, err
	}
	return zones, nil
}

func (g *Client) GetMachineTypes() ([]v1.MachineType, error) {
	zoneList, err := g.GetZones()
	if err != nil {
		return nil, err
	}
	//machinesZone to keep zone to corresponding machine
	machinesZone := map[string][]string{}
	machineTypes := []v1.MachineType{}
	for _, zone := range zoneList {
		req := g.ComputeService.MachineTypes.List(g.GceProjectID, zone)
		err := req.Pages(g.Ctx, func(list *compute.MachineTypeList) error {
			for _, machine := range list.Items {
				res, err := ParseMachine(machine)
				if err != nil {
					return err
				}
				// to check whether we added this machine to machineTypes
				// if we found it then add this zone to machinesZone, else add the machine to machineTypes and also add this zone to machinesZone
				if zones, found := machinesZone[res.Spec.SKU]; found {
					machinesZone[res.Spec.SKU] = append(zones, zone)
				} else {
					machinesZone[res.Spec.SKU] = []string{zone}
					machineTypes = append(machineTypes, *res)
				}
			}
			return nil
		})
		if err != nil {
			return nil, err
		}
	}
	//update g.Data.MachineTypes[].Zones
	for index, instanceType := range machineTypes {
		machineTypes[index].Spec.Zones = machinesZone[instanceType.Spec.SKU]
	}
	return machineTypes, nil
}

func getComputeService(ctx context.Context, credentialFilePath string) (*compute.Service, error) {
	gceInfo := credential.NewGCE()
	err := gceInfo.Load(credentialFilePath)
	if err != nil {
		return nil, err
	}
	conf, err := google.JWTConfigFromJSON([]byte(gceInfo.ServiceAccount()),
		compute.ComputeScope)
	if err != nil {
		return nil, err
	}
	client := conf.Client(ctx)
	return compute.New(client)
}