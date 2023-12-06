package main

import (
	"fmt"
)

func main() {
	//pattern := "abbac"
	//pattern_list := []string{"dog", "cat", "cat", "dog", "fish"}
	//fmt.Println(find_pattern(pattern, pattern_list))
	//word := "cd"
	//list := []string{"ad", "adce", "cda", "ae", "cdb"}
	//fmt.Println(find_letter(word, list))
	//t := "acrr"
	//s := "car"
	//fmt.Println(random_letter(s, t))
	stock := []int{4, 3, 2}
	fmt.Println(stock_list(stock))
}

func find_pattern(pattern_string string, pattern_list []string) bool {
	string_list := []byte(pattern_string)
	result_string := make(map[byte]string)
	count := 0
	for i, value := range pattern_list {
		if count == 0 {
			result_string[string_list[i]] = value
			count = count + 1
			continue
		}
		_, exists := result_string[string_list[i]]
		if exists == true {
			if fmt.Sprint(result_string[string_list[i]]) == value {
				continue
			} else {
				return false
			}
		} else {
			for _, value2 := range result_string {
				if value == fmt.Sprint(value2) {
					return false
				} else {
					result_string[string_list[i]] = value
					break
				}
			}

		}
	}
	return true
}

func find_letter(liense string, letters []string) string {
	var final_word string
	var map_slice []map[byte]int
	count := 0
	for _, value := range letters {
		map_slice = append(map_slice, letters_to_map(value))
	}
	for i, dic := range map_slice {
		if letter_in_map(dic, liense) == true {
			if count == 0 {
				final_word = letters[i]
				count = count + 1
				continue
			}
			if len([]byte(final_word)) > len([]byte(letters[i])) {
				final_word = letters[i]
			}
		}
	}
	return final_word
}

func letters_to_map(letter string) map[byte]int {
	return_map := make(map[byte]int)
	list := []byte(letter)
	for _, value := range list {
		return_map[value] = return_map[value] + 1
	}
	return return_map
}

func letter_in_map(dic map[byte]int, letters string) bool {
	letter_list := []byte(letters)
	for _, v := range letter_list {
		exists := false
		for key := range dic {
			if key == v {
				exists = true
				break
			}
		}
		if exists == true {
			continue
		} else {
			return false
		}

	}
	return true
}

func random_letter(s string, t string) string {
	s_map := letters_to_map(s)
	t_map := letters_to_map(t)
	for k, v := range t_map {
		_, exists := s_map[k]

		if !exists {
			return string(k)
		}

		if exists && v > s_map[k] {
			return string(k)
		}

	}
	return ""
}

func stock_list(stock []int) int {
	if len(stock) <= 1 {
		return 0
	}
	profit := 0
	for buy, buy_value := range stock {
		for i := 0; i < len(stock)-buy; i++ {
			if stock[i+buy] <= 0 || buy_value <= 0 {
				continue
			}
			if buy_value-stock[i+buy] >= profit {
				profit = buy_value - stock[i+buy]
			}
		}
	}
	if profit <= 0 {
		return 0
	}
	return profit
}

func telephone_keypad(number string) []string {
	var final_list []string
	byte_list := []byte(number)
	keypad := map[string]string{
		"2": "abc",
		"3": "def",
		"4": "ghi",
		"5": "jkl",
		"6": "mno",
		"7": "pqrs",
		"8": "tuv",
		"9": "wxyz",
	}
	if len(byte_list) == 1 {
		for _, v := range []byte(keypad[string(byte_list[0])]) {
			final_list = append(final_list, string(v))
		}
		return final_list
	}
	if len(byte_list) > 1 {
		for i, v := range final_list {
			v = v + keypad[string(i)]
		}
		return final_list
	}
	telephone_keypad(string(byte_list[1:]))
}
