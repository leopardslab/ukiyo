package process

import (
	"github.com/olekukonko/tablewriter"
	"github.com/patrickmn/go-cache"
	"os"
	"os/exec"
	"strconv"
	"time"
	"ukiyo/internal/caching/cacheupdate"
	"ukiyo/internal/containerscheduler"
)

type DataObj struct {
	Name         string
	ImageName    string
	ScheduleType string
	TimeMnt      string
	Ports        string
}

func Process(c *cache.Cache) string {
	for {
		var dataObjList []DataObj
		var data [][]string

		pods := containerscheduler.QueryListRecodeInDB()
		if len(pods) > 0 {
			for _, pod := range pods {
				var dataObj DataObj
				dataObj.Name = pod.Name
				if pod.ScheduledDowntime {
					dataObj.ScheduleType = "Waiting For Scheduled"
					a, b, c := cacheupdate.TimeCalculator(pod.Name, c)
					dataObj.ImageName = c
					if a > 0 || b > 0 {
						dataObj.ScheduleType = "Scheduled"
					}
					dataObj.TimeMnt = strconv.Itoa(a) + " [mm]: " + strconv.Itoa(b) + " [SS]"
				} else {
					dataObj.ScheduleType = "ON THE FLY"
					dataObj.ImageName = "-"
					dataObj.TimeMnt = "00 [mm]: " + "00 [SS]"
				}

				if len(pod.BindingPort) > 0 {
					for _, port := range pod.BindingPort {
						dataObj.Ports = dataObj.Ports + port.InternalPort + "->" + port.ExportPort + ", "
					}
				}
				dataObjList = append(dataObjList, dataObj)
			}
		}

		if len(dataObjList) > 0 {
			for _, dataObj := range dataObjList {
				data = append(data, []string{dataObj.Name, dataObj.ImageName, dataObj.ScheduleType, dataObj.TimeMnt, dataObj.Ports})
			}
		}

		table := tablewriter.NewWriter(os.Stdout)
		table.SetRowLine(true)
		table.SetHeader([]string{"Name", "Version", "Deployment Status", "Time Remaining to Deployment", "Ports"})
		table.SetCaption(true, "Caching Table in Ukiyo")

		for _, v := range data {
			table.Append(v)
		}
		table.Render()

		table2 := tablewriter.NewWriter(os.Stdout)
		table2.SetRowLine(true)
		table2.SetHeader([]string{"Time", "Action", "Status", "Comment"})
		table2.SetCaption(true, "History Table")

		for _, v := range cacheupdate.HistoryFetch(c) {
			table2.Append(v)
		}
		table2.Render()

		time.Sleep(1 * time.Second)
		exec := exec.Command("clear")
		exec.Stdout = os.Stdout
		exec.Run()
	}
}
