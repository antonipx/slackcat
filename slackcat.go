/*
 * Simple SlackCat implementation in GoLang using nlopes/slack library
 * Uploads file to a slack channel using API token from env variable
 * v1.0 as@portworx.com
 * 
 */

package main

import (
    "fmt"
    "os"
    
    "github.com/nlopes/slack"
)

func main() {
    if len(os.Args) < 3 {
        fmt.Println("usage: slackcat <channelname> <filename>")
        os.Exit(1)
    }

    sc := slack.New(os.Getenv("APIKEY"))

    channel := make([]string, 1)
    
    conv_params := slack.GetConversationsParameters{
        Limit: 1000,
        Types: []string{"public_channel", "private_channel"},
    }
    
    conv, _, err := sc.GetConversations(&conv_params)
    if err != nil {
        fmt.Println(err)
        return
    }
    for _, c := range conv {
        if c.Name == os.Args[1] {
            channel[0]=c.ID
            break
        }
    }
    
    if(len(channel[0]) != 9) {
        fmt.Println("Channel", os.Args[1], "not found [", channel[0], "]")
        os.Exit(1)
    }
    
    file_params := slack.FileUploadParameters{
        File: os.Args[2],
        Channels:  channel,
    }

    file, err := sc.UploadFile(file_params)
    if err != nil {
        fmt.Println(err, file)
        return
    }

}
