package unit_tests

import (
	"fmt"
	"github.com/atrico-go/tree/display"
	"testing"

	. "github.com/atrico-go/testing/assert"
	"github.com/atrico-go/testing/is"

	"github.com/atrico-go/tree"
)

type displayFuncDef func(config treedisplay.DisplayTreeConfig) []string
type testConfig struct {
	displayFunc   displayFuncDef
	displayConfig treedisplay.DisplayTreeConfig
}
func newTestConfig() testConfig {
	return testConfig{nil, treedisplay.NewDisplayConfig()}
}
func (orig testConfig) ForStandardTree(root tree.Node) testConfig {
	return testConfig{displayFunc: func(config treedisplay.DisplayTreeConfig) []string {
		return treedisplay.DisplayTree(root, config)
	}, displayConfig: orig.displayConfig}
}
func (orig testConfig) ForBinaryTree(root tree.BinaryNode) testConfig {
	return testConfig{displayFunc: func(config treedisplay.DisplayTreeConfig) []string {
		return treedisplay.DisplayBinaryTree(root, config)
	}, displayConfig: orig.displayConfig}
}
func (orig testConfig) WithDisplayType(value treedisplay.DisplayType) testConfig {
	return testConfig{displayFunc: orig.displayFunc, displayConfig: orig.displayConfig.WithDisplayType(value)}
}

func (orig testConfig) WithDisplayCharacterType(value treedisplay.CharacterType) testConfig {
	return testConfig{displayFunc: orig.displayFunc, displayConfig: orig.displayConfig.WithCharacterType(value)}
}

func (orig testConfig) WithHighlight(value interface{}) testConfig {
	return testConfig{displayFunc: orig.displayFunc, displayConfig: orig.displayConfig.WithHighlight(value)}
}

func testDisplayTypesImpl(t *testing.T, testConfig testConfig, expectedTopDown []string, expectedBalanced []string, expectedBalancedFavourTop []string, expectedBottomUp []string) {
	t.Run("Top down", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(treedisplay.TopDown), expectedTopDown) })
	t.Run("Balanced", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(treedisplay.Balanced), expectedBalanced) })
	t.Run("Balanced favour top", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(treedisplay.BalancedFavourTop), expectedBalancedFavourTop) })
	t.Run("Bottom up", func(t *testing.T) { testImpl(t, testConfig.WithDisplayType(treedisplay.BottomUp), expectedBottomUp) })
}

func testImpl(t *testing.T, testConfig testConfig, expected []string) {
	// Act
	result := testConfig.displayFunc(testConfig.displayConfig)
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

