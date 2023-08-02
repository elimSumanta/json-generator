package main

import (
	"encoding/csv"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"os"
	"strconv"
	"strings"

	"github.com/elimSumanta/json-generator/modul"
)

type ProductHighlight struct {
	Title         string `json:"title"`
	GroupOrdering int    `json:"group_order"`
}

type BannerTambahan struct {
	Name  string               `json:"name,omitempty"`
	Ratio string               `json:"ratio,omitempty"`
	Title string               `json:"title,omitempty"`
	Data  []BannerTambahanData `json:"data,omitempty"`
}

type BannerTambahanData struct {
	ImageID  string `json:"image_id,omitempty"`
	ImageURL string `json:"image_url,omitempty"`
	VideoURL string `json:"video_url,omitempty"`
	LinkID   int    `json:"link_id,omitempty"`
	LinkType string `json:"link_type,omitempty"`
}

type LayoutRequest struct {
	ID               string
	ShopName         string
	ShopID           string
	DummyShopID      string
	BGColor          []string
	BGImage          string
	IsDark           bool
	TextColor        string
	HeaderBGColor    string
	ProductHighlight []ProductHighlight
	BannerTambahan   []BannerTambahan
	Kuppon           string
}

func main() {
	// open file
	f, err := os.Open("request.csv")
	if err != nil {
		log.Fatal(err)
	}

	// remember to close the file at the end of the program
	defer f.Close()

	// read csv values using csv.Reader
	csvReader := csv.NewReader(f)
	line := 0
	var req []LayoutRequest
	for {
		rec, err := csvReader.Read()
		if err == io.EOF {
			break
		}
		if err != nil {
			log.Fatal(err)
		}

		if line > 0 {
			// do something with read line
			temp := LayoutRequest{
				ID:          rec[0],
				ShopName:    rec[1],
				ShopID:      rec[2],
				DummyShopID: rec[3],
				BGColor:     strings.Split(rec[4], ","),
				BGImage:     rec[5],
				// IsDark:      strconv.ParseBool(rec[6]),
				TextColor:     rec[7],
				HeaderBGColor: rec[8],
				// BannerTambahan: ,
				Kuppon: rec[11],
			}
			temp.IsDark, err = strconv.ParseBool(rec[6])
			if err != nil {
				log.Printf("Error parsing ID: %s, Row: %+v", rec[0], rec[6])
				continue
			}

			productHighlight := []ProductHighlight{}
			err = json.Unmarshal([]byte(rec[9]), &productHighlight)
			if err != nil {
				log.Printf("Error Unmarshal ID: %s, Row: %+v", rec[0], rec[9])
				continue
			}
			temp.ProductHighlight = productHighlight

			banner := []BannerTambahan{}
			// data := strings.ReplaceAll(rec[10], " ", "")
			data := strings.ReplaceAll(rec[10], "\n", "")
			err = json.Unmarshal([]byte(data), &banner)
			if err != nil {
				log.Printf("Error Unmarshal ID: %s, Row: %+v err:%v", rec[0], data, err.Error())
				continue
			}
			temp.BannerTambahan = banner

			req = append(req, temp)
		}
		line++
	}

	// log.Println(req)
	result := make(map[string]string, 0)
	for _, v := range req {
		widgetList := []string{}
		productHighlightCount := 0
		bannerTambahanCount := 0
		for i, widget := range modul.WidgetOrder {
			switch widget {
			case "banner_timer":
				widgetList = append(widgetList, fmt.Sprintf(modul.BannerWidget, i+1))
			case "product_highlight":
				if productHighlightCount >= len(v.ProductHighlight) {
					widgetList = append(widgetList, fmt.Sprintf(modul.ProductHighlightWidget, i+1, "", productHighlightCount+1))
					continue
				} else {
					widgetList = append(widgetList, fmt.Sprintf(modul.ProductHighlightWidget, i+1, v.ProductHighlight[productHighlightCount].Title, v.ProductHighlight[productHighlightCount].GroupOrdering))
				}
				productHighlightCount++
			case "voucher":
				widgetList = append(widgetList, fmt.Sprintf(modul.Voucher, i+1))
			case "banner_tambahan":
				if bannerTambahanCount >= len(v.BannerTambahan) {
					widgetList = append(widgetList, fmt.Sprintf(modul.InactiveBannerTambahan, i+1))
					continue
				}
				widgetMasterID, ok := modul.MapWidgetMasterID[v.BannerTambahan[bannerTambahanCount].Name]
				if ok {
					data := v.BannerTambahan[bannerTambahanCount].Data
					dataStr := []string{}
					for _, v := range data {
						dataStr = append(dataStr, fmt.Sprintf(modul.BannerTambahanData, v.LinkType, v.LinkID, v.VideoURL, v.ImageID, v.ImageURL, v.ImageURL))
					}
					widgetList = append(widgetList, fmt.Sprintf(modul.BannerTambahan, widgetMasterID, i+1, v.BannerTambahan[bannerTambahanCount].Name, v.BannerTambahan[bannerTambahanCount].Ratio, v.BannerTambahan[bannerTambahanCount].Title, 1, strings.Join(dataStr, ",")))
				}
				bannerTambahanCount++
			case "play":
				widgetList = append(widgetList, fmt.Sprintf(modul.Play, i+1))
			case "slider_banner_highlight":
				widgetList = append(widgetList, fmt.Sprintf(modul.SliderBanner, i+1))
			}
		}

		bgColor := `\"` + strings.Join(v.BGColor, `\",\"`) + `\"`
		widgetStr := strings.Join(widgetList, ",")
		templ := fmt.Sprintf(modul.Template, v.ShopID, string(bgColor), v.BGImage, v.TextColor, v.IsDark, v.HeaderBGColor, string(widgetStr))
		result[v.ID] = templ
		// log.Println(result)
	}

	WriteToFile(result)
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func WriteToFile(result map[string]string) {
	for key, v := range result {
		f, err := os.Create("layout/layoutID_" + key)
		check(err)

		n3, err := f.WriteString(v)
		check(err)
		fmt.Printf("wrote %d bytes\n", n3)
		f.Sync()
		f.Close()
	}
}
