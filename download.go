package main

import (
	"bufio"
	"context"
	"fmt"
	"os"
	"strings"

	"github.com/antchfx/htmlquery"
	"github.com/chromedp/chromedp"
)

func Download(xs XzInfo) {
	page, _ := htmlquery.LoadURL(xs.FirstUrl)
	xcontent := fmt.Sprintf(`%s`, xs.Content)
	content := htmlquery.FindOne(page, xcontent)
	title := htmlquery.FindOne(page, fmt.Sprintf(`%s`, xs.Title))
	nextpage := htmlquery.FindOne(page, fmt.Sprintf(`%s`, xs.NextPage))
	if content == nil || title == nil || nextpage == nil {
		fmt.Println("用了DP")
		xs.xsxzdp()
	} else {
		fmt.Println("没用")
		xs.xsxzh()
	}

}

func (xs *XzInfo) xsxzdp() {
	file, _ := os.OpenFile(xs.Name+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)

	defer file.Close()
	writer := bufio.NewWriter(file)
	ctx, cancel := chromedp.NewContext(context.Background())
	defer cancel()
	page_url := xs.FirstUrl
	var content, title, nextPage string
	var ok bool
	for {

		err := chromedp.Run(ctx,
			chromedp.Navigate(page_url),
			chromedp.Text(xs.Title, &title, chromedp.NodeVisible),
			chromedp.Text(xs.Content, &content, chromedp.NodeVisible),
			chromedp.AttributeValue(xs.NextPage, "href", &nextPage, &ok, chromedp.NodeVisible),
		)
		if err != nil {
			fmt.Println(err)
		}
		writer.Write([]byte(strings.ReplaceAll(title+"\n"+content+"\n", "<br>", "\\r\\n")))

		if nextPage[len(nextPage)-4:] != xs.BreakFlag {
			break
		} else {
			page_url = xs.HeadUrl + nextPage
		}

	}

	writer.Flush()
}

func (xs *XzInfo) xsxzh() {
	file, _ := os.OpenFile(xs.Name+".txt", os.O_RDWR|os.O_APPEND|os.O_CREATE, 0664)
	cycleFlag := 0
	defer file.Close()
	writer := bufio.NewWriter(file)
	var c, t string
	page_url := xs.FirstUrl
	for {
		if cycleFlag > 3 {
			break
		}

		page, _ := htmlquery.LoadURL(page_url)
		xcontent := fmt.Sprintf(`%s`, xs.Content)
		content := htmlquery.FindOne(page, xcontent)
		title := htmlquery.FindOne(page, fmt.Sprintf(`%s`, xs.Title))
		nextpage := htmlquery.FindOne(page, fmt.Sprintf(`%s`, xs.NextPage))
		if content == nil {
			cycleFlag++
			continue
		} else {
			cycleFlag = 0
		}
		c = htmlquery.InnerText(content)
		t = htmlquery.InnerText(title)
		writer.Write([]byte(strings.ReplaceAll(t+"\n"+c+"\n", "<br>", "\\r\\n")))
		nexturl := htmlquery.SelectAttr(nextpage, "href")
		if nexturl[len(nexturl)-4:] != xs.BreakFlag {
			break
		} else {
			page_url = xs.HeadUrl + htmlquery.SelectAttr(nextpage, "href")
		}

	}
	writer.Flush()
}
