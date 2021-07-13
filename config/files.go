// B"H
package config

type Config struct {
	ExtentionsIDs []string `json:"ExtentionsIDs"`
	QueryUrl      string   `json:"queryUrl"`
}

// type Config struct {
// 	Files     []File     `json:"files"`
// 	Endpoints []Endpoint `json:"endpoints"`
// }
// type Urls struct {
// 	Windows string `json:"windows"`
// 	Linux   string `json:"linux"`
// }
// type Files struct {
// 	// Name string `json:"name"`
//   // Urls Urls   `json:"urls"`
//   Files []File `json:"files"`
// }
// type Endpoint struct {
// 	Name string `json:"name"`
// 	URL  string `json:"url"`
// }

// type Files struct {
//   Files []File `json:"files"`
// }

// // File struct which contains a name
// // and a list of files links
// type File struct {
// 	Name      string `json:"name"`
// 	Urls      Urls   `json:"urls"`
// 	Extention string `json:"extention,omitempty"`
// }

// type Config struct {
//   Clusters Endpoints //`json:"endpoints"`
//   Binaries Files //`json:"files"`
// }

// type Endpoints struct {
//   Endpoints []Endpoint `json:"endpoints"`
// }

// type Endpoint struct {
//   Name string `json:"name"`
//   Url  string `json:"url"`
// }
// // Files struct which contains
// // an array of files
// type Files struct {
//   Files []File `json:"files"`
// }

// // File struct which contains a name
// // and a list of files links
// type File struct {
//   Name string `json:"name"`
//   Urls Urls   `json:"urls"`
//   Extention string `json:"extention,omitempty"`
// }

// // Urls struct which contains a
// // list of links
// type Urls struct {
//   Windows string `json:"windows"`
//   Linux   string `json:"linux"`
// }
