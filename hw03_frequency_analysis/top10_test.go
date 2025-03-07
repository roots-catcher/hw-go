package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

var text2 = `Какая-то рандомная строка для теста,которую я написал и хочу оттестить. 
			Тест теста оттестить тест. Для тестирования топ десяти слов?`

func Test2Top10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("positive test", func(t *testing.T) {
		expected := []string{
			"Для",
			"Какая-то",
			"Тест",
			"десяти",
			"для",
			"и",
			"написал",
			"оттестить",
			"оттестить.",
			"рандомная",
		}
		require.Equal(t, expected, Top10(text2))
	})
}
