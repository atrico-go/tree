package unit_tests

import (
	"fmt"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/tree"
)

func testTypesImpl(t *testing.T, displayFunc func(config tree.DisplayTreeConfig) []string, expectedTopDown []string, expectedBalanced []string, expectedBalancedFavourTop []string, expectedBottomUp []string) {
	t.Run("Top down", func(t *testing.T) { testImpl(t, displayFunc, tree.TopDown, expectedTopDown) })
	t.Run("Balanced", func(t *testing.T) { testImpl(t, displayFunc, tree.Balanced, expectedBalanced) })
	t.Run("Balanced favour top", func(t *testing.T) { testImpl(t, displayFunc, tree.BalancedFavourTop, expectedBalancedFavourTop) })
	t.Run("Bottom up", func(t *testing.T) { testImpl(t, displayFunc, tree.BottomUp, expectedBottomUp) })
}

func testImpl(t *testing.T, displayFunc func(config tree.DisplayTreeConfig) []string, displayType tree.DisplayType, expected []string) {
	// Arrange
	config := tree.DisplayTreeConfig{Type: displayType}

	// Act
	result := displayFunc(config)
	display(result)

	// Assert
	assertDisplay(t, result, expected)
}

func display(lines []string) {
	for _, ln := range lines {
		fmt.Println(ln)
	}
}

func assertDisplay(t *testing.T, actual []string, expected []string) {
	Assert(t).That(len(actual), is.EqualTo(len(expected)), "Correct number of lines")
	for i := range actual {
		Assert(t).That(actual[i], is.EqualTo(expected[i]), fmt.Sprintf("Line %d", i))
	}
}
