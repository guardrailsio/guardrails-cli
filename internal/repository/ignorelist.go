package repository

import "strings"

var ignoredFileTypes = []string{
	".aac", ".aiff", ".ape", ".au", ".flac", ".gsm", ".it", ".m3u", ".m4a", ".mid", ".mod", ".mp3", ".mpa", ".pls", ".ra", ".s3m", ".sid", ".wav", ".wma", ".xm", ".7z",
	".a", ".ar", ".bz2", ".cab", ".cpio", ".deb", ".dmg", ".egg", ".gz", ".iso", ".lha", ".mar", ".pea", ".rar", ".rpm", ".s7z", ".shar", ".tar", ".tbz2", ".tgz", ".tlz",
	".whl", ".xpi", ".deb", ".rpm", ".xz", ".pak", ".crx", ".exe", ".msi", ".bin", ".eot", ".otf", ".ttf", ".woff", ".woff2", ".3dm", ".3ds", ".max", ".bmp", ".dds", ".gif",
	".jpg", ".jpeg", ".png", ".psd", ".xcf", ".tga", ".thm", ".tif", ".tiff", ".yuv", ".ai", ".eps", ".ps", ".svg", ".dwg", ".dxf", ".gpx", ".kml", ".kmz", ".ods", ".xls",
	".xlsx", ".csv", ".ics", ".vcf", ".ppt", ".odp", ".3g2", ".3gp", ".aaf", ".asf", ".avchd", ".avi", ".drc", ".flv", ".m2v", ".m4p", ".m4v", ".mkv", ".mng", ".mov", ".mp2",
	".mp4", ".mpe", ".mpeg", ".mpg", ".mpv", ".mxf", ".nsv", ".ogg", ".ogv", ".ogm", ".qt", ".rm", ".rmvb", ".roq", ".srt", ".svi", ".vob", ".webm", ".wmv", ".yuv",
}

func isFileTypeIgnored(filename string) bool {
	var isIgnored bool
	for _, ignoredFileType := range ignoredFileTypes {
		if strings.HasSuffix(filename, ignoredFileType) {
			isIgnored = true
		}
	}

	return isIgnored
}
