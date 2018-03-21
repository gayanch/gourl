package analytics

type Analytics interface {
    Get(shortcode string) string
    Update(shortcode string)
}

func GetAnalyticsService() Analytics {
    return InMemoryAnalytics{data: make(map[string]string)}
}

// const (
//     ANALYTICS_DIR = "analytics/"
// )

// func init() {
//     os.Mkdir(ANALYTICS_DIR, os.ModePerm)
// }

// func Get(shortcode string) string {
//     if file, err := os.Open(ANALYTICS_DIR + shortcode); err == nil {
//         defer file.Close()

//         var ret string
//         fmt.Fscanf(file, "%s", &ret)

//         return ret
//     } else {
//         return "No Analytics Found for " + shortcode
//     }
// }

// func Update(shortcode string) {
//     if file, err := os.Open(ANALYTICS_DIR + shortcode); err == nil {
//         defer file.Close()

//         var count int
//         fmt.Fscanf(file, "%d", &count)

//         count++
//         fmt.Println(count)
//         fmt.Fprintf(file, "%d", count)
//     } else {
//         newfile, _ := os.Create(ANALYTICS_DIR + shortcode)
//         defer newfile.Close()

//         fmt.Fprintf(newfile, "%d", 0)
//     }
// }
