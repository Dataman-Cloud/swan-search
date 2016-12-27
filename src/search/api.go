package search

import (
	"errors"
	"net/http"
	"sort"

	"github.com/gin-gonic/gin"
	"github.com/renstrom/fuzzysearch/fuzzy"
)

const (
	RESULT_LEN = 10
)

const (
	//Search
	CodeInvalidSearchKeywords = "13001"
	CodeOk                    = "0"
)

func (searchApi *SearchApi) Search(ctx *gin.Context) {
	query := ctx.Query("keyword")
	if query == "" {
		err := errors.New("invalid keyword")
		ctx.JSON(http.StatusBadRequest, gin.H{"code": CodeInvalidSearchKeywords, "message": err.Error()})
		return
	}

	results := []*Document{}
	indexs := fuzzy.RankFind(query, searchApi.Index)
	sort.Sort(indexs)
	if len(indexs) > 0 {
		if len(indexs) > 10 {
			indexs = indexs[:RESULT_LEN]
		}
		for _, index := range indexs {
			results = append(results, searchApi.Store.Get(index.Target))
		}
	}
	ctx.JSON(http.StatusOK, gin.H{"code": CodeOk, "data": results})
	return
}
