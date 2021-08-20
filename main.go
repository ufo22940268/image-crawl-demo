// Command screenshot is a chromedp example demonstrating how to take a
// screenshot of a specific element and of the entire browser viewport.
package main

import (
	"context"
	"flag"
	"fmt"
	"github.com/chromedp/cdproto/network"
	"io/ioutil"
	"log"
	"os"

	"github.com/chromedp/chromedp"
)

func main() {
	fmt.Println("----------------startj---------------------")
	// create context
	devtoolsWsURL := flag.String("devtools-ws-url", "ws://127.0.0.1:9222", "ws://127.0.0.1:9222")
	flag.Parse()
	fmt.Printf("devtools url: %s \n", *devtoolsWsURL)
	if *devtoolsWsURL == "" {
		log.Fatal("must specify -devtools-ws-url")
	}

	// create allocator context for use with creating a browser context later
	allocatorContext, cancel := chromedp.NewRemoteAllocator(context.Background(), *devtoolsWsURL)
	ctx, cancel := chromedp.NewContext(allocatorContext)
	defer cancel()

	// capture screenshot of an element
	var buf []byte

	// capture entire browser viewport, returning png with quality=90
	if err := chromedp.Run(ctx, fullScreenshot(`https://staging.compass.com/app/ads-center/digital/get-ad-photo?adType=4&originalUrl=https%3A%2F%2Fstaging.compass.com%2Fm%2Fb371e9f8b2771caf77cf5184a497f6f8b41c5bff_img_0%2Forigin.jpg&adFocusChoiceVal=listing&bannerTitle=5225%20Pooks%20Hill%20Road%2C%20Unit%20616N&bannerSubtitle=JUST%20LISTED&bannerSize=Bethesda%20%7C%202%20Bed%2C%201%20Bath&transform=%7B%22scale%22%3A1.143%2C%22translate%22%3A%5B0%2C0%5D%7D&aspect=square`, 90, &buf)); err != nil {
		log.Fatal(err)
	}
	filename := "fullScreenshot.png"
	os.Remove(filename)
	if err := ioutil.WriteFile(filename, buf, 0o644); err != nil {
		log.Fatal(err)
	}

	log.Printf("wrote fullScreenshot.png")
}

// fullScreenshot takes a screenshot of the entire browser viewport.
//
// Note: chromedp.FullScreenshot overrides the device's emulation settings. Reset
func fullScreenshot(urlstr string, quality int, res *[]byte) chromedp.Tasks {
	return chromedp.Tasks{
		network.SetExtraHTTPHeaders(map[string]interface{}{
			"cookie": os.Getenv("cookie"),
		}),
		chromedp.Navigate(urlstr),
		chromedp.FullScreenshot(res, quality),
	}
}
