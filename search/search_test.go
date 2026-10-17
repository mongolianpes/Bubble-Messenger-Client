package search

import (
	"fmt"
	"strings"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestSearchUser(t *testing.T) {
	loginPiece := generateTestRandValidStr()
	login1, name1, password1, device1, key1, err := RegAndAuthTestUser(loginPiece)
	require.NoError(t, err)
	login2, name2, _, _, _, err := RegAndAuthTestUser(loginPiece)
	require.NoError(t, err)
	login3, name3, password3, device3, key3, err := RegAndAuthTestUser(generateTestRandValidStr())
	require.NoError(t, err)
	fmt.Println(login1)
	fmt.Println(login2)
	fmt.Println(login3)

	searchxactMatch := SearchUser(login3, password3, device3, "@"+login1, key3, 1)
	searchxactMatchSplitted := strings.Split(searchxactMatch, " ")
	require.Equal(t, login1, searchxactMatchSplitted[0])
	require.Equal(t, name1, searchxactMatchSplitted[1])

	searchxactMatch = SearchUser(login1, password1, device1, "@"+login2, key1, 1)
	searchxactMatchSplitted = strings.Split(searchxactMatch, " ")
	require.Equal(t, login2, searchxactMatchSplitted[0])
	require.Equal(t, name2, searchxactMatchSplitted[1])

	searchxactMatch = SearchUser(login1, password1, device1, "@"+login3, key1, 1)
	searchxactMatchSplitted = strings.Split(searchxactMatch, " ")
	require.Equal(t, login3, searchxactMatchSplitted[0])
	require.Equal(t, name3, searchxactMatchSplitted[1])

	search := SearchUser(login3, password3, device3, loginPiece, key3, 1)
	searchSplitted := strings.Split(search, "\n")
	for _, line := range searchSplitted {
		lineSplitted := strings.Split(line, " ")
		require.Contains(t, []string{login1, login2}, lineSplitted[0])
		require.Contains(t, []string{name1, name2}, lineSplitted[1])
	}
}
