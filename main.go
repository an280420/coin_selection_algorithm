package main

import (
	"bufio" // буферизация
	"fmt"
	"os"      // пакет для чтения ввода с консоли
	"strconv" // пакет для конвертации строковых значений в числа
	"strings" // пакет для конвертации строки в массив слов
)

func main() {
	var money int
	var nominalsCount int
	fmt.Println("Привествую пользователь. \nДанная утилита представляет возможность получить наименьшее количество монет различного номинала, необходимых для формирования некой суммы 'money'")
	fmt.Println("Программа принимает значения входных параметров в следующем формате \n3 \n1 3 4 \n10 \nгде первая строка - натуральное число - количество номиналов монет \nвторая строка - перечисление номиналов монет \nтретья строка - натуральное число - сумма, которую необходимо получить из монет доступного номинала")
	fmt.Println("Введите значения входных параметров:")
	fmt.Scanf("%d\n", &nominalsCount)      // примем из консоли число номиналов монет (избыточный аргумент, требующий проверки)
	coinsString := scanLineWithSpaces()    // примем из консоли строку с пробелами
	coins := stringToIntArray(coinsString) // преобразуем строку в массив монет (проверка на уникальность не предумострена)
	fmt.Scanf("%d\n", &money)              // примем из консоли сумму которую необходимо получить

	// валидация входных данных
	if money > 10000 {
		fmt.Println("Параметр в третье строке должен быть от 0 до 10000. Введите данный параметр снова")
		for money > 10000 {
			fmt.Scanf("%d\n", &money)
		}
	}
	if nominalsCount != len(coins) {
		fmt.Println("Внимание значение первой строки не соответствует числу номиналов указанных во второй строке.\nКоличество номиналов монет взято из второй строки")
	}

	// вычисление количества монет
	fmt.Println("Минимальное количество монет для этого ")
	printChange(coins, money)
}

//выбор номиналов
// Эта функция берет массив монет и сумму которую надо набрать. Возвращает часть выбора от coins.
func change(coins []int, money int) []int {
	// инициализация массива выбора монет
	// в массиве будем хранить порядковые номера доступных монет (порядоковый номер не по индексу, а именно порядковый номер)
	// таким образом чтобы в каждой позиции массива будем записывать ту монету которую необходимо взять в первую очереь чтобы начать набирать сумму соответствующую индексу в массиве
	// пример: монеты 3, 5, первая монета 3, вторая - 5. требуемая сумма 7
	// массив который будет получатся [0, 0, 0, 1, 0, 2, 1, 0] (для суммы 3 - необходима монета по номером 1, для суммы 5 - монета по номером 2, для суммы 6 - монета 1)
	// длина массива будет соответствовать числу которым представленна итоговая сумма
	S := make([]int, money+1)
	// вспомогательный массив
	C := make([]int, money+1)
	C[0] = 0

	// решение от малого к большему (от нулевой суммы до требуемой), m - текущая позиция (сумма)
	for m := 1; m <= money; m++ {
		C[m] = 10001 // максимальный номинал

		// для каждой монеты набора проверить можно ли вернуть
		for j := 0; j < len(coins); j++ {
			if m >= coins[j] && C[m-coins[j]]+1 < C[m] {
				C[m] = C[m-coins[j]] + 1 // сумма порядоквых номеров монет
				S[m] = j + 1             // сохранение порядокового номера монеты для текущей суммы - m
			}
		}
	}

	return S
}

// получает набор монет (массивом) и сумму которую необходиом набрать этим набором
// набирает монеты для набора money (идем от money к 0)
func printChange(coins []int, money int) {
	S := change(coins, money)

	coinsCountForMoney := 0

	for money > 0 {
		if S[money]-1 < 0 {
			fmt.Println("IMPOSSIBLE")
			break
		} else {
			// coins[S[money]-1] 		выбранная монета для набора
			coinsCountForMoney++
			money = money - coins[S[money]-1]
		}
	}

	if coinsCountForMoney > 0 {
		fmt.Println(coinsCountForMoney)
	}
}

// преобразование строки в массив чисел
func stringToIntArray(someString string) []int {
	wordsArray := strings.Fields(someString)

	integerArray := []int{}

	for _, i := range wordsArray {
		j, err := strconv.Atoi(i)
		if err != nil {
			panic(err)
		}
		integerArray = append(integerArray, j)
	}
	return integerArray
}

// функция сканирования текста с пробелами
func scanLineWithSpaces() string {
	myscanner := bufio.NewScanner(os.Stdin)
	myscanner.Scan()
	line := myscanner.Text()
	return line
}
