package main

import (
	"os"
	"fmt"
	"time"
	"bufio"
	"strings"
	"crypto/sha1"
	"encoding/hex"
)


/*
Wrench
Date: 08/29/21
Author: 0x1CA3
*/


func clear_screen() {
	fmt.Println("\033[H\033[2J")
}

func display_banner() {
	fmt.Println(` 
		_      __________  ____  _____/ /_ 
   		| | /| / / ___/ _ \/ __ \/ ___/ __ \
   		| |/ |/ / /  /  __/ / / / /__/ / / /
   		|__/|__/_/   \___/_/ /_/\___/_/ /_/ 
   		[Made by: https://github.com/0x1CA3]
   		Usage: ./wrench <hash> <wordlist.txt>
   `)
}

func crack_hash(line string, hash string) {
	process_hash := sha1.New()
	process_hash.Write([]byte(line))
	
	str_hash := hex.EncodeToString(process_hash.Sum(nil))
	if hash == str_hash {
		file_name := "password_found.txt"
		open_file, err := os.OpenFile(file_name, os.O_RDWR|os.O_APPEND|os.O_CREATE, 0660)
		if err != nil {
			os.Exit(1)
		} else {
			defer open_file.Close()
			fmt.Fprintf(open_file, "\n  		Hash Cracked! Password -> %s\n", line)
		}
	}
	return
}

func read_wordlist(targ_file string, hash string) {
	passwords, err := os.Open(targ_file)
	if err != nil {
		fmt.Println("[!] Error, could not open/find wordlist file! [!]")
	} else {
		defer passwords.Close()
		scanr := bufio.NewScanner(passwords)
		for scanr.Scan() {
			line := scanr.Text()
			crack_hash(line, hash)
		}
		cracked_password_ouput, err := os.Open("password_found.txt")
		if err != nil {
			fmt.Println("\n[!] Could not crack hash! [!]")
		} else {
			defer cracked_password_ouput.Close()
			cscanr := bufio.NewScanner(cracked_password_ouput)
			for cscanr.Scan() {
				fmt.Println(cscanr.Text())
			}
		}
	}
	return
}

func hash_compare(hash string) {
	database := map[interface{}]string {
		1: "d033e22ae348aeb5660fc2140aec35850c4da997",
		2: "5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8",
		3: "e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4",
		4: "9d4e1e23bd5b727046a9e3b4b7db57bd8d6ee684",
		5: "7c4a8d09ca3762af61e59520943dc26494f8941b",
		6: "d8be2ef007bd17d46b3cd126861cd58655a40c33",
		7: "7c222fb2927d828af22f592134e8932480637c0d",
		8: "b1b3773a05c0ed0176787a4f1574ff0075f7521e",
		9: "f7c3bc1d808e04732adf679965ccc34ca7ae3441",
		10: "8cb2237d0679ca88db6464eac60da96345513964",
		11: "7110eda4d09e062aa5e4a390b0a572ac0d2c0220",
		12: "3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d",
		13: "20eabe5d64b0e216796e834f52d61fd0b70332fc",
		14: "af8978b1797b72acfff9595a5a2a373ec3d9106d",
		15: "601f1889667efaebb33b8c12572835da3f027f78",
		16: "a2c901c8c6dea98958c219f6f2d038c44dc5d362",
		17: "6367c48dd193d56ea7b0baad25b19455e529f5ee",
		18: "2d27b62c597ec858f6e7b54e7e58525e6a95e6d8",
		19: "ab87d24bdc7452e55738deb5f868e1f16dea5ace",
		20: "b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3",
		21: "cedf41fccb586dc39e1ce34bb482f0afe557b49f",
		22: "ed9d3d832af899035363a69fd53cd3be8f71501c",
		23: "4f26aeafdb2367620a393c973eddbe8f8b846ebd",
		24: "1411678a0b9e25ee2f7c8b2f7ac92b6a74b3f9c5",
		25: "b0399d2029f64d445bd131ffaa399a42d2f8e7dc",
		26: "4d9012b4a77a9524d675dad27c3276ab5705e5e8",
		27: "40123e9c6273385ea69892c48c80aa6cb25b9113",
		28: "01b307acba4f54f55aafc33bb06bbbf6ca803e9a",
		29: "17b9e1c64588c7fa6419b4d29dc1f4426279ba01",
		30: "dd5fef9c1c1da1394d6d34b248c51be2ad740840",
		31: "35ed5406781ebfdf7161bbbb18e16cb9ad1f3be4",
		32: "18c28604dd31094a8d69dae60f1bcd347f1afc5a",
		33: "c6922b6ba9e0939583f973bc1682493351ad4fe8",
		34: "74a871acbf060dda5fc7260d05a5924a34e4c0e7",
		35: "dd2edb87ea9eb7a32fd4057276d3a1fab861c1d5",
		36: "48058e0c99bf7d689ce71c360699a14ce2f99774",
		37: "c984aed014aec7623a54f0591da07a85fd4b762d",
		38: "cb45c671cbc500627ea424eea5f91996221b5935",
		39: "05fe7461c607c33229772d402505601016a7d0ea",
		40: "59033478180d07080d5e4f3baa0099996c364162",
		41: "e68e11be8b70e435c65aef8ba9798ff7775c361e",
		42: "1cb5bd5a9e45420321f44c72da5d90d7f0432ffb",
		43: "e3cd9f6469fc3e1acfb9f2bdbfc5a3d2bbb8e2ad",
		44: "93ec71b22793a81569c94ca17e4d9c293d8e201f",
		45: "7ab515d12bd2cf431745511ac4ee13fed15ab578",
		46: "6e2f9e6111e77edd0c446ea7a84e25323d137a61",
		47: "1999e4893f732ba38b948dbe8d34ed48cd54f058",
		48: "5c17fa03e6d5fc247565e1cd8ffa70e1bfe5b8d9",
		49: "f32157a45887e4fe5adc0b5198f7ec4920a526d7",
		50: "5c6d9edc3a951cda763f650235cfc41a3fc23fe8",
	}

	pass_database := map[string]string {
		"d033e22ae348aeb5660fc2140aec35850c4da997": "admin",
		"5baa61e4c9b93f3f0682250b6cf8331b7ee68fd8": "password",
		"e5e9fa1ba31ecd1ae84f75caaa474f3a663f05f4": "secret",
		"9d4e1e23bd5b727046a9e3b4b7db57bd8d6ee684": "pass",
		"7c4a8d09ca3762af61e59520943dc26494f8941b": "123456",
		"d8be2ef007bd17d46b3cd126861cd58655a40c33": "password÷NO PASSWORDS!",
		"7c222fb2927d828af22f592134e8932480637c0d": "12345678",
		"b1b3773a05c0ed0176787a4f1574ff0075f7521e": "qwerty",
		"f7c3bc1d808e04732adf679965ccc34ca7ae3441": "123456789",
		"8cb2237d0679ca88db6464eac60da96345513964": "12345",
		"7110eda4d09e062aa5e4a390b0a572ac0d2c0220": "1234",
		"3d4f2bf07dc1be38b20cd6e46949a1071f9d0e3d": "111111",
		"20eabe5d64b0e216796e834f52d61fd0b70332fc": "1234567",
		"af8978b1797b72acfff9595a5a2a373ec3d9106d": "dragon",
		"601f1889667efaebb33b8c12572835da3f027f78": "123123",
		"a2c901c8c6dea98958c219f6f2d038c44dc5d362": "baseball",
		"6367c48dd193d56ea7b0baad25b19455e529f5ee": "abc123",
		"2d27b62c597ec858f6e7b54e7e58525e6a95e6d8": "football",
		"ab87d24bdc7452e55738deb5f868e1f16dea5ace": "monkey",
		"b7a875fc1ea228b9061041b7cec4bd3c52ab3ce3": "letmein",
		"cedf41fccb586dc39e1ce34bb482f0afe557b49f": "696969",
		"ed9d3d832af899035363a69fd53cd3be8f71501c": "shadow",
		"4f26aeafdb2367620a393c973eddbe8f8b846ebd": "master",
		"1411678a0b9e25ee2f7c8b2f7ac92b6a74b3f9c5": "666666",
		"b0399d2029f64d445bd131ffaa399a42d2f8e7dc": "qwertyuiop",
		"4d9012b4a77a9524d675dad27c3276ab5705e5e8": "123321",
		"40123e9c6273385ea69892c48c80aa6cb25b9113": "mustang",
		"01b307acba4f54f55aafc33bb06bbbf6ca803e9a": "1234567890",
		"17b9e1c64588c7fa6419b4d29dc1f4426279ba01": "michael",
		"dd5fef9c1c1da1394d6d34b248c51be2ad740840": "654321",
		"35ed5406781ebfdf7161bbbb18e16cb9ad1f3be4": "pussy",
		"18c28604dd31094a8d69dae60f1bcd347f1afc5a": "superman",
		"c6922b6ba9e0939583f973bc1682493351ad4fe8": "1qaz2wsx",
		"74a871acbf060dda5fc7260d05a5924a34e4c0e7": "7777777",
		"dd2edb87ea9eb7a32fd4057276d3a1fab861c1d5": "fuckyou",
		"48058e0c99bf7d689ce71c360699a14ce2f99774": "121212",
		"c984aed014aec7623a54f0591da07a85fd4b762d": "000000",
		"cb45c671cbc500627ea424eea5f91996221b5935": "qazwsx",
		"05fe7461c607c33229772d402505601016a7d0ea": "123qwe",
		"59033478180d07080d5e4f3baa0099996c364162": "killer",
		"e68e11be8b70e435c65aef8ba9798ff7775c361e": "trustno1",
		"1cb5bd5a9e45420321f44c72da5d90d7f0432ffb": "jordan",
		"e3cd9f6469fc3e1acfb9f2bdbfc5a3d2bbb8e2ad": "jennifer",
		"93ec71b22793a81569c94ca17e4d9c293d8e201f": "zxcvbnm",
		"7ab515d12bd2cf431745511ac4ee13fed15ab578": "asdfgh",
		"6e2f9e6111e77edd0c446ea7a84e25323d137a61": "hunter",
		"1999e4893f732ba38b948dbe8d34ed48cd54f058": "buster",
		"5c17fa03e6d5fc247565e1cd8ffa70e1bfe5b8d9": "soccer",
		"f32157a45887e4fe5adc0b5198f7ec4920a526d7": "harley",
		"5c6d9edc3a951cda763f650235cfc41a3fc23fe8": "batman",
	}

	time.Sleep(3 * time.Second)
	
	for i := 0; i <= 50; i++ {
		if hash == database[i] {
			fmt.Println("\n  		Password Found! ->", pass_database[hash])
		}
	}
	return
}

func validation_check(hash string) {
	var char_check int = 0
	hash_length := len(hash)	
	
	if hash_length == 40 {
		fmt.Println("		Test 1 -> Passed!")
	} else {
		fmt.Println("		Test 1 -> Failed!")
	}

	valid_db := map[interface{}]string {
		1: "a",
		2: "b",
		3: "c",
		4: "d",
		5: "e",
		6: "f",
		7: "1",
		8: "2",
		9: "3",
		10: "4",
		11: "5",
		12: "6",
		13: "7",
		14: "8",
		15: "9",
		16: "0",
	}

	for i := 0; i <= 16; i++ {
		if strings.Contains(hash, valid_db[i]) {
			char_check++
		}
	}
	if char_check == 16 || char_check == 17 {
		fmt.Println("		Test 2 -> Passed!")
	} else {
		fmt.Println("		Test 2 -> Failed!")
	}
	hash_compare(hash)
}

func main() {
	clear_screen()
	display_banner()
	
	clean_f := os.Remove("password_found.txt")
	if clean_f != nil {}
	
	main_arg := os.Args[1]
	targ_file := os.Args[2]
	hash_bytes := []byte(main_arg)

	fmt.Println("		Target Hash -> ", main_arg)
	fmt.Println("                Extra -> ", hash_bytes)
	
	fmt.Println(`		
		[Validation check]
	    	------------------`)
	
	validation_check(main_arg)
	read_wordlist(targ_file, main_arg)
}