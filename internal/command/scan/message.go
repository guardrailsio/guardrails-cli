package scan

import (
	"fmt"

	"github.com/enescakir/emoji"
	"github.com/jedib0t/go-pretty/text"
)

func (h *Handler) stopLoadingMessage() {
	if !h.Args.Quiet {
		h.Spinner.Stop()
	}
}

func (h *Handler) displayScanningMessage() {
	if !h.Args.Quiet {
		fmt.Println()
		h.Spinner.Prefix = text.FgCyan.Sprintf("%s scanning %s ", emoji.Eyes, h.Args.Path)
		h.Spinner.FinalMSG = text.FgCyan.Sprintf("%s scanning %s %s\n\n", emoji.Eyes, h.Args.Path, emoji.CheckMark)
		h.Spinner.Start()
	}
}

func (h *Handler) displayCompressingMessage(projectName string) {
	if !h.Args.Quiet {
		fmt.Println()
		h.Spinner.Prefix = text.FgCyan.Sprintf("%s compressing %s ", emoji.Books, projectName)
		h.Spinner.FinalMSG = text.FgCyan.Sprintf("%s compressing %s %s\n", emoji.Books, projectName, emoji.CheckMark)
		h.Spinner.Start()
	}
}

func (h *Handler) displayUploadingMessage(projectZipName string) {
	if !h.Args.Quiet {
		fmt.Println()
		h.Spinner.Prefix = text.FgCyan.Sprintf("%s uploading %s ", emoji.Package, projectZipName)
		h.Spinner.FinalMSG = text.FgCyan.Sprintf("%s uploading %s %s\n", emoji.Package, projectZipName, emoji.CheckMark)
		h.Spinner.Start()
	}
}

func (h *Handler) displayRetrievingScanResultMessage(projectName string) {
	if !h.Args.Quiet {
		fmt.Println()
		h.Spinner.Prefix = text.FgCyan.Sprintf("%s retrieving scan result of %s ", emoji.MagnifyingGlassTiltedRight, projectName)
		h.Spinner.FinalMSG = text.FgCyan.Sprintf("%s retrieving scan result of %s %s\n\n", emoji.MagnifyingGlassTiltedRight, projectName, emoji.CheckMark)
		h.Spinner.Start()
	}
}
