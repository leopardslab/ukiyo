package process

import (
	"github.com/olekukonko/tablewriter"
	"github.com/patrickmn/go-cache"
	"os"
	"os/exec"
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
					dataObj.ScheduleType = "true"
					a, b, c := cacheupdate.TimeCalculator(pod.Name, c)
					dataObj.ImageName = c
					dataObj.TimeMnt = a + " [mm]: " + b + " [SS]"
				} else {
					dataObj.ScheduleType = "false"
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
		table.SetHeader([]string{"Name", "Image Name", "Build Scheduled", "Time", "Exec Command"})

		for _, v := range data {
			table.Append(v)
		}
		table.Render()

		time.Sleep(1 * time.Second)
		c := exec.Command("clear")
		c.Stdout = os.Stdout
		c.Run()
	}
}
