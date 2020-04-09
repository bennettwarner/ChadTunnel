package main

import (
	"fmt"
	"github.com/AlecAivazis/survey/v2"
	"github.com/AlecAivazis/survey/v2/terminal"
	"github.com/fatih/color"
	"github.com/phayes/freeport"
	"log"
	"os"
	"os/exec"
	"os/signal"
	"os/user"
	"strconv"
	"strings"
	"syscall"
)

func execInput(input string) error {
	// Remove the newline character.
	input = strings.TrimSuffix(input, "\n")

	// Split the input separate the command and the arguments.
	args := strings.Split(input, " ")

	// Check for built-in commands.
	switch args[0] {
	case "cd":
		// 'cd' to home with empty path not yet supported.
		if len(args) < 2 {
			return os.Chdir("~")
		}
		// Change the directory and return the error.
		return os.Chdir(args[1])
	case "exit":
		os.Exit(0)
	}

	// Prepare the command to execute.
	cmd := exec.Command(args[0], args[1:]...)

	// Set the correct output device.
	cmd.Stdin  = os.Stdin
	cmd.Stderr = os.Stderr
	cmd.Stdout = os.Stdout

	// Execute the command and return the error.
	return cmd.Run()
}

func SetupCloseHandler() {
	c := make(chan os.Signal, 2)
	signal.Notify(c, os.Interrupt, syscall.SIGTERM)
	go func() {
		<-c
		fmt.Println("\r- Ctrl+C pressed in Terminal")
		os.Exit(0)
	}()
}

func proxy() {
	localUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	jumpbox := false
	prompt := &survey.Confirm{
		Message: "Do you want to connect to this remote host via a jumpbox?",
	}
	err1 := survey.AskOne(prompt, &jumpbox)
	if err1 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err1 != nil {
		panic(err1)
	}
	if jumpbox{
		jumpbox_ip := ""
		prompt1 := &survey.Input{
			Message: "Jumpbox IP/Hostname:",
		}
		err := survey.AskOne(prompt1, &jumpbox_ip, survey.WithValidator(survey.Required))
		if err == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err != nil {
			panic(err)
		}

		jumpbox_port := ""
		prompt6 := &survey.Input{
			Message: "Jumpbox SSH Port:",
			Default: "22",
		}
		err2 := survey.AskOne(prompt6, &jumpbox_port, survey.WithValidator(survey.Required))
		if err2 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err2 != nil {
			panic(err2)
		}

		jumpbox_username := ""
		prompt2 := &survey.Input{
			Message: "Jumpbox Username:",
			Default: localUser.Username,
		}
		err3 := survey.AskOne(prompt2, &jumpbox_username, survey.WithValidator(survey.Required))
		if err3 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err3 != nil {
			panic(err3)
		}

		rhost_ip := ""
		prompt3 := &survey.Input{
			Message: "Remote Host IP/Hostname:",
		}
		err4 := survey.AskOne(prompt3, &rhost_ip, survey.WithValidator(survey.Required))
		if err4 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err4 != nil {
			panic(err4)
		}

		rhost_port := ""
		prompt5 := &survey.Input{
			Message: "Remote Host SSH Port:",
			Default: "22",
		}
		err5 := survey.AskOne(prompt5, &rhost_port, survey.WithValidator(survey.Required))
		if err5 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err5 != nil {
			panic(err5)
		}

		rhost_username := ""
		prompt4 := &survey.Input{
			Message: "Remote Host Username:",
			Default: localUser.Username,
		}
		err6 := survey.AskOne(prompt4, &rhost_username, survey.WithValidator(survey.Required))
		if err6 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err6 != nil {
			panic(err6)
		}

		proxy_port := ""
		prompt7 := &survey.Input{
			Message: "Proxy Port:",
			Default: strconv.Itoa(port),
		}
		err7 := survey.AskOne(prompt7, &proxy_port, survey.WithValidator(survey.Required))
		if err7 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err7 != nil {
			panic(err7)
		}

		command := "ssh -D " + proxy_port + " -N -J " + jumpbox_username + "@" + jumpbox_ip + ":" + jumpbox_port + " " + rhost_username + "@" + rhost_ip + " -p " + rhost_port
		color.Green("\n[+] Launching SOCKS5 proxy on port " + proxy_port + "...\n")
		fmt.Println("    Press Ctrl + C to exit")

		fmt.Println("\nIn the future you can save time and launch the command directly:")
		color.Magenta("    " + command)
		execInput(command)
	} else {
		rhost_ip := ""
		prompt1 := &survey.Input{
			Message: "Remote Host IP/Hostname:",
		}
		err := survey.AskOne(prompt1, &rhost_ip, survey.WithValidator(survey.Required))
		if err == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err != nil {
			panic(err)
		}

		rhost_port := ""
		prompt2 := &survey.Input{
			Message: "Remote Host SSH Port:",
			Default: "22",
		}
		err1 := survey.AskOne(prompt2, &rhost_port, survey.WithValidator(survey.Required))
		if err1 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err1 != nil {
			panic(err1)
		}

		rhost_username := ""
		prompt3 := &survey.Input{
			Message: "Remote Host Username:",
			Default: localUser.Username,
		}
		err2 := survey.AskOne(prompt3, &rhost_username, survey.WithValidator(survey.Required))
		if err2 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err2 != nil {
			panic(err2)
		}

		proxy_port := ""
		prompt7 := &survey.Input{
			Message: "Proxy Port:",
			Default: strconv.Itoa(port),
		}
		err7 := survey.AskOne(prompt7, &proxy_port, survey.WithValidator(survey.Required))
		if err7 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err7 != nil {
			panic(err7)
		}
		command := "ssh -D " + proxy_port + " -N " + rhost_username + "@" + rhost_ip + " -p " + rhost_port
		color.Green("\n[+] Launching SOCKS5 proxy on port " + proxy_port + "...\n")
		fmt.Println("    Press Ctrl + C to exit\n")

		fmt.Println("In the future you can save time and launch the command directly:")
		color.Magenta("    " + command)
		execInput(command)
	}
}

func portForward() {
	localUser, err := user.Current()
	if err != nil {
		panic(err)
	}
	port, err := freeport.GetFreePort()
	if err != nil {
		log.Fatal(err)
	}
	jumpbox := false
	prompt := &survey.Confirm{
		Message: "Do you want to connect to this remote host via a jumpbox?",
	}
	err1 := survey.AskOne(prompt, &jumpbox)
	if err1 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err1 != nil {
		panic(err1)
	}
	if jumpbox{
		jumpbox_ip := ""
		prompt1 := &survey.Input{
			Message: "Jumpbox IP/Hostname:",
		}
		err := survey.AskOne(prompt1, &jumpbox_ip, survey.WithValidator(survey.Required))
		if err == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err != nil {
			panic(err)
		}

		jumpbox_port := ""
		prompt6 := &survey.Input{
			Message: "Jumpbox SSH Port:",
			Default: "22",
		}
		err2 := survey.AskOne(prompt6, &jumpbox_port, survey.WithValidator(survey.Required))
		if err2 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err2 != nil {
			panic(err2)
		}

		jumpbox_username := ""
		prompt2 := &survey.Input{
			Message: "Jumpbox Username:",
			Default: localUser.Username,
		}
		err3 := survey.AskOne(prompt2, &jumpbox_username, survey.WithValidator(survey.Required))
		if err3 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err3 != nil {
			panic(err3)
		}

		pivot_ip := ""
		prompt3 := &survey.Input{
			Message: "Pivot Host IP/Hostname:",
		}
		err4 := survey.AskOne(prompt3, &pivot_ip, survey.WithValidator(survey.Required))
		if err4 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err4 != nil {
			panic(err4)
		}

		pivot_port := ""
		prompt5 := &survey.Input{
			Message: "Pivot Host SSH Port:",
			Default: "22",
		}
		err5 := survey.AskOne(prompt5, &pivot_port, survey.WithValidator(survey.Required))
		if err5 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err5 != nil {
			panic(err5)
		}

		pivot_username := ""
		prompt4 := &survey.Input{
			Message: "Pivot Host Username:",
			Default: localUser.Username,
		}
		err6 := survey.AskOne(prompt4, &pivot_username, survey.WithValidator(survey.Required))
		if err6 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err6 != nil {
			panic(err6)
		}

		rhost_ip := ""
		prompt7 := &survey.Input{
			Message: "Remote Host IP/Hostname:",
		}
		err7 := survey.AskOne(prompt7, &rhost_ip, survey.WithValidator(survey.Required))
		if err7 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")
			os.Exit(0)
		} else if err7 != nil {
			panic(err7)
		}

		rhost_port := ""
		prompt8 := &survey.Input{
			Message: "Remote Host Port:",
		}
		err8 := survey.AskOne(prompt8, &rhost_port, survey.WithValidator(survey.Required))
		if err8 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")
			os.Exit(0)
		} else if err8 != nil {
			panic(err8)
		}

		command := "ssh -L " + strconv.Itoa(port) + ":" + rhost_ip + ":" + rhost_port + " -N -J " + jumpbox_username + "@" + jumpbox_ip + ":" + jumpbox_port + " " + pivot_username + "@" + pivot_ip + " -p " + pivot_port
		color.Green("\n[+] Launching SSH Port Forwarding...")
		fmt.Println("    You can now access the remote service at localhost:" + strconv.Itoa(port))
		fmt.Println("    Press Ctrl + C to exit")

		fmt.Println("\nIn the future you can save time and launch the command directly:")
		color.Magenta("    " + command)
		execInput(command)
	} else {
		pivot_ip := ""
		prompt3 := &survey.Input{
			Message: "Pivot Host IP/Hostname:",
		}
		err4 := survey.AskOne(prompt3, &pivot_ip, survey.WithValidator(survey.Required))
		if err4 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err4 != nil {
			panic(err4)
		}

		pivot_port := ""
		prompt5 := &survey.Input{
			Message: "Pivot Host SSH Port:",
			Default: "22",
		}
		err5 := survey.AskOne(prompt5, &pivot_port, survey.WithValidator(survey.Required))
		if err5 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err5 != nil {
			panic(err5)
		}

		pivot_username := ""
		prompt4 := &survey.Input{
			Message: "Pivot Host Username:",
			Default: localUser.Username,
		}
		err6 := survey.AskOne(prompt4, &pivot_username, survey.WithValidator(survey.Required))
		if err6 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")

			os.Exit(0)
		} else if err6 != nil {
			panic(err6)
		}

		rhost_ip := ""
		prompt7 := &survey.Input{
			Message: "Remote Host IP/Hostname:",
		}
		err7 := survey.AskOne(prompt7, &rhost_ip, survey.WithValidator(survey.Required))
		if err7 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")
			os.Exit(0)
		} else if err7 != nil {
			panic(err7)
		}

		rhost_port := ""
		prompt8 := &survey.Input{
			Message: "Remote Host Port:",
		}
		err8 := survey.AskOne(prompt8, &rhost_port, survey.WithValidator(survey.Required))
		if err8 == terminal.InterruptErr {
			fmt.Println("\r- Ctrl+C pressed in Terminal")
			os.Exit(0)
		} else if err8 != nil {
			panic(err8)
		}

		command := "ssh -L " + strconv.Itoa(port) + ":" + rhost_ip + ":" + rhost_port + " -N " + pivot_username + "@" + pivot_ip + " -p " + pivot_port
		color.Green("\n[+] Launching SSH Port Forwarding...")
		fmt.Println("    You can now access the remote service at localhost:" + strconv.Itoa(port))
		fmt.Println("    Press Ctrl + C to exit")

		fmt.Println("\nIn the future you can save time and launch the command directly:")
		color.Magenta("    " + command)
		execInput(command)
	}
}

func ssh() {
	localUser, err := user.Current()
	if err != nil {
		panic(err)
	}

	jumpbox_ip := ""
	prompt1 := &survey.Input{
		Message: "Jumpbox IP/Hostname:",
	}
	err8 := survey.AskOne(prompt1, &jumpbox_ip, survey.WithValidator(survey.Required))
	if err8 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err8 != nil {
		panic(err8)
	}

	jumpbox_port := ""
	prompt6 := &survey.Input{
		Message: "Jumpbox SSH Port:",
		Default: "22",
	}
	err2 := survey.AskOne(prompt6, &jumpbox_port, survey.WithValidator(survey.Required))
	if err2 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err2 != nil {
		panic(err2)
	}

	jumpbox_username := ""
	prompt2 := &survey.Input{
		Message: "Jumpbox Username:",
		Default: localUser.Username,
	}
	err3 := survey.AskOne(prompt2, &jumpbox_username, survey.WithValidator(survey.Required))
	if err3 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err3 != nil {
		panic(err3)
	}

	rhost_ip := ""
	prompt3 := &survey.Input{
		Message: "Remote Host IP/Hostname:",
	}
	err4 := survey.AskOne(prompt3, &rhost_ip, survey.WithValidator(survey.Required))
	if err4 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err4 != nil {
		panic(err4)
	}

	rhost_port := ""
	prompt5 := &survey.Input{
		Message: "Remote Host SSH Port:",
		Default: "22",
	}
	err5 := survey.AskOne(prompt5, &rhost_port, survey.WithValidator(survey.Required))
	if err5 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err5 != nil {
		panic(err5)
	}

	rhost_username := ""
	prompt4 := &survey.Input{
		Message: "Remote Host Username:",
		Default: localUser.Username,
	}
	err6 := survey.AskOne(prompt4, &rhost_username, survey.WithValidator(survey.Required))
	if err6 == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err6 != nil {
		panic(err6)
	}

	command := "ssh " + "-J " + jumpbox_username + "@" + jumpbox_ip + ":" + jumpbox_port + " " + rhost_username + "@" + rhost_ip + " -p " + rhost_port
	color.Green("\n[+] Launching interactive SSH session on\n    " + rhost_ip + " via " + jumpbox_ip + "	\n")
	fmt.Println("\n    Press Ctrl + C to exit")

	fmt.Println("\nIn the future you can save time and launch the command directly:")
	color.Magenta("    " + command)
	execInput(command)
}

func main() {
	SetupCloseHandler()

	title := `
                                            ,#####,
                                            #_   _#
                                            |a' 'a|
 ██████╗██╗  ██╗ █████╗ ██████╗             |  u  |
██╔════╝██║  ██║██╔══██╗██╔══██╗            \  =  /
██║     ███████║███████║██║  ██║            |\___/|
██║     ██╔══██║██╔══██║██║  ██║   ___ ____/:     :\____ ___
╚██████╗██║  ██║██║  ██║██████╔╝ .'   '.-===-\   /-===-.'   '.
 ╚═════╝╚═╝  ╚═╝╚═╝  ╚═╝╚═════╝ /      .-"""""-.-"""""-.      \
████████╗██╗   ██╗███╗   ██╗███╗   ██╗███████╗██╗
╚══██╔══╝██║   ██║████╗  ██║████╗  ██║██╔════╝██║
   ██║   ██║   ██║██╔██╗ ██║██╔██╗ ██║█████╗  ██║
   ██║   ██║   ██║██║╚██╗██║██║╚██╗██║██╔══╝  ██║
   ██║   ╚██████╔╝██║ ╚████║██║ ╚████║███████╗███████╗
   ╚═╝    ╚═════╝ ╚═╝  ╚═══╝╚═╝  ╚═══╝╚══════╝╚══════╝
`
	color.Red(title)
	color.Green("\tVersion 2003 | Written by @b_e_n & tnkr\n\n\n")

	var mode int
	prompt := &survey.Select{
		Message: "Choose an action:",
		Options: []string{"I want to proxy traffic through a host", "I want to access a remote service through a host", "I want SSH access to a remote host via a jumpbox"},
	}
	err := survey.AskOne(prompt, &mode)
	if err == terminal.InterruptErr {
		fmt.Println("\r- Ctrl+C pressed in Terminal")

		os.Exit(0)
	} else if err != nil {
		panic(err)
	}

	switch mode{
	case 0:
		proxy()
	case 1:
		portForward()
	case 2:
		ssh()
	}
}