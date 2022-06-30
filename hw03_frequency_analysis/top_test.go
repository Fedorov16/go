package hw03frequencyanalysis

import (
	"testing"

	"github.com/stretchr/testify/require"
)

// Change to true if needed.
var taskWithAsteriskIsCompleted = true

var text = `Как видите, он  спускается  по  лестнице  вслед  за  своим
	другом   Кристофером   Робином,   головой   вниз,  пересчитывая
	ступеньки собственным затылком:  бум-бум-бум.  Другого  способа
	сходить  с  лестницы  он  пока  не  знает.  Иногда ему, правда,
		кажется, что можно бы найти какой-то другой способ, если бы  он
	только   мог   на  минутку  перестать  бумкать  и  как  следует
	сосредоточиться. Но увы - сосредоточиться-то ему и некогда.
		Как бы то ни было, вот он уже спустился  и  готов  с  вами
	познакомиться.
	- Винни-Пух. Очень приятно!
		Вас,  вероятно,  удивляет, почему его так странно зовут, а
	если вы знаете английский, то вы удивитесь еще больше.
		Это необыкновенное имя подарил ему Кристофер  Робин.  Надо
	вам  сказать,  что  когда-то Кристофер Робин был знаком с одним
	лебедем на пруду, которого он звал Пухом. Для лебедя  это  было
	очень   подходящее  имя,  потому  что  если  ты  зовешь  лебедя
	громко: "Пу-ух! Пу-ух!"- а он  не  откликается,  то  ты  всегда
	можешь  сделать вид, что ты просто понарошку стрелял; а если ты
	звал его тихо, то все подумают, что ты  просто  подул  себе  на
	нос.  Лебедь  потом  куда-то делся, а имя осталось, и Кристофер
	Робин решил отдать его своему медвежонку, чтобы оно не  пропало
	зря.
		А  Винни - так звали самую лучшую, самую добрую медведицу
	в  зоологическом  саду,  которую  очень-очень  любил  Кристофер
	Робин.  А  она  очень-очень  любила  его. Ее ли назвали Винни в
	честь Пуха, или Пуха назвали в ее честь - теперь уже никто  не
	знает,  даже папа Кристофера Робина. Когда-то он знал, а теперь
	забыл.
		Словом, теперь мишку зовут Винни-Пух, и вы знаете почему.
		Иногда Винни-Пух любит вечерком во что-нибудь поиграть,  а
	иногда,  особенно  когда  папа  дома,  он больше любит тихонько
	посидеть у огня и послушать какую-нибудь интересную сказку.
		В этот вечер...`

var textPotter = `
	It was on the corner of the street that he noticed the first sign 
	of something peculiar - a cat reading a map. For a second, Mr 
	Dursley didn’t realise what he had seen - then he jerked his head 
	around to look again. There was a tabby cat standing on the corner 
	of Privet Drive, but there wasn’t a map in sight. What could 
	he have been thinking of? It must have been a trick of the light. 
	Mr Dursley blinked and stared at the cat. It stared back. As Mr 
	Dursley drove around the corner and up the road, he watched the 
	cat in his mirror. It was now reading the sign that said Privet Drive 
	- no, looking at the sign; cats couldn’t read maps or signs. Mr 
	Dursley gave himself a little shake and put the cat out of his 
	mind. As he drove towards town he thought of nothing except a 
	large order of drills he was hoping to get that day.
	`

var textPotterLessText = `
	It was on the corner
`

var textMostCommonMistakes = `
	Нога нога! нога, 'нога' какой-то какойто
	`

func TestTop10(t *testing.T) {
	t.Run("no words in empty string", func(t *testing.T) {
		require.Len(t, Top10(""), 0)
	})

	t.Run("Vini Puh", func(t *testing.T) {
		if taskWithAsteriskIsCompleted {
			expected := []string{
				"а",         // 8
				"он",        // 8
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"в",         // 4
				"его",       // 4
				"если",      // 4
				"кристофер", // 4
				"не",        // 4
			}
			require.Equal(t, expected, Top10(text))
		} else {
			expected := []string{
				"он",        // 8
				"а",         // 6
				"и",         // 6
				"ты",        // 5
				"что",       // 5
				"-",         // 4
				"Кристофер", // 4
				"если",      // 4
				"не",        // 4
				"то",        // 4
			}
			require.Equal(t, expected, Top10(text))
		}
	})
	t.Run("Harry Potter", func(t *testing.T) {
		expected := []string{
			"the",     // 16
			"a",       // 8
			"he",      // 8
			"of",      // 7
			"cat",     // 5
			"dursley", // 4
			"it",      // 4
			"mr",      // 4
			"was",     // 4
			"and",     // 3
		}
		require.Equal(t, expected, Top10(textPotter))
	})
	t.Run("Harry Potter less words", func(t *testing.T) {
		require.Len(t, Top10(textPotterLessText), 5)
	})
	t.Run("Most common mistakes", func(t *testing.T) {
		expected := []string{
			"нога",     // 4
			"какой-то", // 1
			"какойто",  // 1
		}
		require.Equal(t, expected, Top10(textMostCommonMistakes))
	})
}
