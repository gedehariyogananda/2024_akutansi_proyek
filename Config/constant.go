package Config

var AllowedImageExtensions = map[string]bool{
	".jpg":  true,
	".jpeg": true,
	".png":  true,
}

var MaxMultipartMemory = 10 << 20 // default is 10 MB
var StaticFileDir = "./public/storage"
var StaticFileRoute = "/storage"
