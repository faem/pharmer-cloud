package options

import (
	"strings"

	"github.com/pharmer/pharmer/utils"
	"github.com/spf13/cobra"
	"github.com/spf13/pflag"
)

type ClusterGetConfig struct {
	Clusters []string
	Output   string
	Owner    string
}

func NewClusterGetConfig() *ClusterGetConfig {
	return &ClusterGetConfig{
		Output: "",
		Owner:  utils.GetLocalOwner(),
	}
}

func (c *ClusterGetConfig) AddFlags(fs *pflag.FlagSet) {
	fs.StringVarP(&c.Output, "output", "o", c.Output, "Output format. One of: json|yaml|wide")
	fs.StringVarP(&c.Owner, "owner", "", c.Owner, "Current user id")
}

func (c *ClusterGetConfig) ValidateFlags(cmd *cobra.Command, args []string) error {
	c.Clusters = func(names []string) []string {
		for i := range names {
			names[i] = strings.ToLower(names[i])
		}
		return names
	}(args)
	return nil
}