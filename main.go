package main

import (
	"github.com/code043/go-cli-pixrem-converter/pixelunit"
	"github.com/code043/go-cli-pixrem-converter/remunit"
	"github.com/spf13/cobra"
)

func main(){
	var rootCmd = &cobra.Command{
		Use: "goc",
	}
	var px, rem string
	var cmd = &cobra.Command{
		Use: "goc",
		Short: "Pixel and rem conversor.",
		Run: func(cmd *cobra.Command, args []string){
			if px != ""{
				pixelunit.PixelToRem(px)				
			}
			if rem != ""{
				remunit.RemToPx(rem)
			}			
		},
	}
	cmd.Flags().StringVarP(&px, "px", "p", "", "Convert rem to pixel")
	cmd.Flags().StringVarP(&rem, "rem", "r", "", "Convert pixel to rem")
	rootCmd.AddCommand(cmd)
	rootCmd.Execute()
}


