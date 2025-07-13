package handlers

import (
	"mmyoungman/templ/utils"
	"mmyoungman/templ/views/pages"
	"net/http"
)

func HandleClickButtonLoadPartial() HTTPHandler {
	return func(w http.ResponseWriter, r *http.Request) error {
		baseArgs := utils.GenerateBaseArgs(r)
		return pages.ExamplesPartial(baseArgs).Render(r.Context(), w)
	}
}

func HandleTest1(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return pages.ExamplePartialTest1().Render(r.Context(), w)
}

func HandleTest2(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return pages.ExamplePartialTest2().Render(r.Context(), w)
}

func HandlePlaceButtonInTarget1(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return pages.ExamplePartialPlaceButtonInTarget1().Render(r.Context(), w)
}

func HandlePlaceButtonInTarget2(w http.ResponseWriter, r *http.Request) error {
	// @MarkFix can visit {URL}/test directly in a browser
	return pages.ExamplePartialPlaceButtonInTarget2().Render(r.Context(), w)
}
