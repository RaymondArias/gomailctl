package cmd

import (
	"fmt"
	"os"
	"strings"

	"github.com/RaymondArias/gomailctl/internal/mailsend"

	"github.com/spf13/cobra"

	homedir "github.com/mitchellh/go-homedir"
	"github.com/spf13/viper"
)

var (
	// mail content
	from      string
	recipient string
	content   string

	//smtp auth
	username string
	password string
	smtpHost string

	cfgFile string
)

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "gomailctl",
	Short: "simple smtp client",
	Long: `A longer description that spans multiple lines and likely contains
examples and usage of using your application. For example:

Cobra is a CLI library for Go that empowers applications.
This application is a tool to generate the needed files
to quickly create a Cobra application.`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Printf("Sending email\n")
		hostData := strings.Split(smtpHost, ":")
		emailData := mailsend.Mail{
			From:       from,
			Recipients: recipient,
			Data:       []byte(content),
			Username:   username,
			Password:   password,
			SMTPServer: hostData[0],
			SMTPPort:   hostData[1],
		}
		mailsend.SendMail(emailData)

	},
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	cobra.OnInitialize(initConfig)

	// Here you will define your flags and configuration settings.
	// Cobra supports persistent flags, which, if defined here,
	// will be global for your application.

	rootCmd.PersistentFlags().StringVar(&cfgFile, "config", "", "config file (default is $HOME/.gomailctl.yaml)")
	rootCmd.PersistentFlags().StringVarP(&from, "from", "f", "", "from address")
	rootCmd.MarkFlagRequired("from")
	rootCmd.PersistentFlags().StringVarP(&recipient, "recipient", "r", "", "recipient")
	rootCmd.MarkFlagRequired("recipient")
	rootCmd.PersistentFlags().StringVarP(&content, "content", "c", "", "content of email")
	rootCmd.MarkFlagRequired("content")
	rootCmd.PersistentFlags().StringVarP(&username, "username", "u", "", "username for smtp auth")
	rootCmd.MarkFlagRequired("username")
	rootCmd.PersistentFlags().StringVarP(&password, "password", "p", "", "password for smtp auth")
	rootCmd.MarkFlagRequired("password")
	rootCmd.PersistentFlags().StringVarP(&smtpHost, "smtp server", "s", "", "smtp server to connect to: hostname:port")
	rootCmd.MarkFlagRequired("smtp server")

	// Cobra also supports local flags, which will only run
	// when this action is called directly.
	rootCmd.Flags().BoolP("toggle", "t", false, "Help message for toggle")
}

// initConfig reads in config file and ENV variables if set.
func initConfig() {
	if cfgFile != "" {
		// Use config file from the flag.
		viper.SetConfigFile(cfgFile)
	} else {
		// Find home directory.
		home, err := homedir.Dir()
		if err != nil {
			fmt.Println(err)
			os.Exit(1)
		}

		// Search config in home directory with name ".gomailctl" (without extension).
		viper.AddConfigPath(home)
		viper.SetConfigName(".gomailctl")
	}

	viper.AutomaticEnv() // read in environment variables that match

	// If a config file is found, read it in.
	if err := viper.ReadInConfig(); err == nil {
		fmt.Println("Using config file:", viper.ConfigFileUsed())
	}
}
