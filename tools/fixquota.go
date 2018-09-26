package main

import (
	"os"

	"github.com/alibaba/pouch/storage/quota"

	"github.com/sirupsen/logrus"
	"github.com/spf13/cobra"
)

var (
	dir       string
	quotaID   uint32
	recursive bool
)

func run(cmd *cobra.Command) error {
	_, err := quota.StartQuotaDriver(dir)
	if err != nil {
		logrus.Errorf("failed to start quota driver for %s, err: %v", dir, err)
		return err
	}

	_, err = quota.SetSubtree(dir, quotaID)
	if err != nil {
		logrus.Errorf("failed to set subtree for %s, quota id: %d, err: %v", dir, quotaID, err)
		return err
	}

	if recursive {
		if err := quota.SetQuotaForDir(dir, quotaID); err != nil {
			logrus.Errorf("failed to set quota id for %s recursively, quota id: %d, err: %v", dir, quotaID, err)
			return err
		}
	}

	return nil
}

func setupFlags(cmd *cobra.Command) {
	flagSet := cmd.Flags()

	flagSet.StringVarP(&dir, "dir", "d", "", "The directory is set quota id.")
	flagSet.Uint32VarP(&quotaID, "quota-id", "i", 0, "The quota id to set.")
	flagSet.BoolVarP(&recursive, "recursive", "r", true, "Set the directory recursively.")
}

func main() {
	var cmdServe = &cobra.Command{
		Use:          "fixquota <-d /path> <-i id>",
		Args:         cobra.NoArgs,
		SilenceUsage: true,
		RunE: func(cmd *cobra.Command, args []string) error {
			return run(cmd)
		},
	}

	setupFlags(cmdServe)
	if err := cmdServe.Execute(); err != nil {
		logrus.Error(err)
		os.Exit(1)
	}
}
