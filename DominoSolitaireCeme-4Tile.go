// DATA KELOMPOK
// JUDUL : DOMINO SOLITAIRE CEME-4TILE
// ANGGOTA 1 : MOCH. NAUVAL RIZALDI NASRIL / 1301194482
// ANGGOTA 2 : JOSHUA ERLANGGA SAKTI / 1301190226
// KELAS : IF-43-08

package main

import (
	"fmt"
	"math/rand"
	"os"
	"os/exec"
	"time"
)

const N = 100000

var player [N]playerinfo
var numplayers int

type domtile struct {
	// consists of the two numbers on the left and right side of the domino tiles and also contain a boolean which marks a double tile
	left, right int
	balak       bool
}

type playerinfo struct {
	// contains all the player's info for the scoreboard
	name        string
	rate        float32
	score, game int
}

func main() {
	// this is where the magic happens
	var domino [28]domtile
	var respondmenu, respondgame, pidx int
	var mrate float32
	var p playerinfo
	var ptile, dtile [4]domtile
	// the header with the title of the game
	fmt.Println("WELCOME TO THE GAME OF DOMINO CEME-4TILE")
	fmt.Println("Created by Joshua and Nauval of Tel-U.")
	fmt.Println("-----------------------------------------------------------------------------------------------------------")
	AnyKey("Press any key and then enter to continue...  ")
	clear()
	// main menu
	for respondmenu != 4 { // 4 means exiting the game
		respondmenu = -1
		fmt.Println("-- DOMINO CEME-4TILE --")
		fmt.Println(" ")
		fmt.Println("***MAIN MENU***")
		fmt.Print("1. Start Playing\n", "2. View The Scoreboard\n", "3. Learn How To Play\n", "4. Exit Game\n")
		fmt.Print("What do you want to do ? (1/2/3/4)\n", "--> ")
		fmt.Scan(&respondmenu)
		for respondmenu < 1 || respondmenu > 4 {
			fmt.Print("Invalid input.\n", "The valid inputs are 1-4.\n", "Try inputting again.\n", "--> ")
			fmt.Scan(&respondmenu)
		}
		fmt.Println(" ")
		if respondmenu == 1 {
			fmt.Println("--------------------------------------------------------------------------------------------------GAMESTART")
			fmt.Println("**LET US START THE GAME**")
			// inputing player's name
			p.score = 0
			p.game = 0
			fmt.Print("Insert your name. Note that names are case-sensitive.\n", "--> ")
			fmt.Scan(&p.name)
			namecheck(&p, &pidx)
			clear()
			fmt.Print("Hello ", p.name, ".\n")
			printscore(p)
			// beginning the game
			createtiles(&domino)
			respondgame = -1
			for respondgame != 9 { // 9 means going back to the main menu
				respondgame = -1
				scramble(domino, &domino)
				spreaddomino(domino, &ptile, &dtile)
				fmt.Println("Dealing...")
				switchdomino(domino, &ptile, &respondgame)
				if respondgame == 0 {
					fmt.Println(" ")
					fmt.Println("----------------------------------------------------------------------------------------------------VERDICT")
					fmt.Println("It's settled!")
					fmt.Println("Your tiles     :", "(", ptile[0].left, "|", ptile[0].right, ")", "-", "(", ptile[1].left, "|", ptile[1].right, ")", "-", "(", ptile[2].left, "|", ptile[2].right, ")", "-", "(", ptile[3].left, "|", ptile[3].right, ")")
					fmt.Println("Dealer's tiles :", "(", dtile[0].left, "|", dtile[0].right, ")", "-", "(", dtile[1].left, "|", dtile[1].right, ")", "-", "(", dtile[2].left, "|", dtile[2].right, ")", "-", "(", dtile[3].left, "|", dtile[3].right, ")")
					gameresult(ptile, dtile, &p)
				} else if respondgame == 5 {
					// if the player want to switch
					mrate = (float32(p.score) / float32(p.game)) * 100
					if p.game == 0 {
						mrate = 0
					}
					player[pidx] = p
					fmt.Println(" ")
					fmt.Print("Game Over for ", p.name, ".\n", "Your winning rate is ", mrate, "%.\n")
					p.game = 0
					p.score = 0
					fmt.Println("-----------------------------------------------------------------------------------------------PLAYERSWITCH")
					fmt.Print("Insert the new player's name. Note that names are case-sensitive.\n", "--> ")
					fmt.Scan(&p.name)
					namecheck(&p, &pidx)
					clear()
					fmt.Print("Hello ", p.name, ".\n")
					printscore(p)
				}
			}
			// the player's result is stored in the array
			mrate = (float32(p.score) / float32(p.game)) * 100
			if p.game == 0 {
				mrate = 0
			}
			player[pidx] = p
			fmt.Print("Game Over.\n", "Your winning rate is ", mrate, "%.\n")
			fmt.Println("---------------------------------------------------------------------------------------------------GAMEOVER")
			AnyKey("Press any key and then enter to go back...  ")
			clear()
		} else if respondmenu == 2 {
			fmt.Println("-------------------------------------------------------------------------------------------------SCOREBOARD")
			showscoreboard()
			clear()
		} else if respondmenu == 3 {
			fmt.Println("--------------------------------------------------------------------------------------------------HOWTOPLAY")
			instructions()
			clear()
		}
	}
	// goodbye message
	fmt.Println("Thank you for playing!")
	AnyKey("Press any key and then enter to exit...  ")
}

// FOR DOMINO

func createtiles(domino *[28]domtile) {
	// a procedure to create the 28 domino tiles
	var a, b, c int

	for a = 0; a <= 6; a++ {
		for b = a; b <= 6; b++ {
			if a == b {
				domino[c] = domtile{a, b, true}
			} else {
				domino[c] = domtile{a, b, false}
			}
			c++
		}
	}
}

func random() {
	// a procedure to create rand seeds
	rand.Seed(time.Now().UnixNano())
}

func printscore(p playerinfo) {
	// a procedure to print "win" and "game" in plural or singular form depending on the situation
	if p.score > 1 && p.game > 1 {
		fmt.Print("Your score : ", p.score, " wins from ", p.game, " games.\n")
	} else if p.score > 1 && p.game <= 1 {
		fmt.Print("Your score : ", p.score, " wins from ", p.game, " game.\n")
	} else if p.score <= 1 && p.game > 1 {
		fmt.Print("Your score : ", p.score, " win from ", p.game, " games.\n")
	} else {
		fmt.Print("Your score : ", p.score, " win from ", p.game, " game.\n")
	}
}

func scramble(dominoIn [28]domtile, dominoOut *[28]domtile) {
	// a procedure to scramble the domino array
	var idx [28]int
	var count, n int
	var valid bool
	for i := 0; i < 28; i++ {
		idx[i] = -1
	}
	for count != 28 {
		random()
		n = rand.Intn(28)
		valid = true
		for i := 0; i < count && valid != false; i++ {
			if idx[i] == n {
				valid = false
			}
		}
		if valid == true {
			idx[count] = n
			dominoOut[n] = dominoIn[count]
			count++
		}
	}
}

func spreaddomino(domino [28]domtile, ptile *[4]domtile, dtile *[4]domtile) {
	// a procedure to spread the domino tiles to the player and the dealer
	var help int
	for i := 0; i < 4; i++ {
		ptile[i] = domino[i]
	}
	for i := 27; i > 23; i-- {
		dtile[help] = domino[i]
		help++
	}

}

func switchdomino(domino [28]domtile, ptile *[4]domtile, respondgame *int) {
	// a procedure to handle if the player wants to change their tiles
	var count int = 0
	var help int = 4
	for count < 2 && *respondgame != 0 && *respondgame != 9 && *respondgame != 5 {
		fmt.Println("Your tiles :", "(", ptile[0].left, "|", ptile[0].right, ")", "-", "(", ptile[1].left, "|", ptile[1].right, ")", "-", "(", ptile[2].left, "|", ptile[2].right, ")", "-", "(", ptile[3].left, "|", ptile[3].right, ")")
		fmt.Print("What's your decision ? (1/2/3/4/5/0/9) (1-4 = Change tile 1-4, 5 = Switch Player, 0 = Done, 9 = Exit)\n", "--> ")
		fmt.Scan(&*respondgame)
		for *respondgame < 0 || *respondgame > 5 && *respondgame < 9 || *respondgame > 9 {
			fmt.Print("Invalid input.\n", "The valid inputs are 0-4 and 9.\n", "Try inputting again.\n", "--> ")
			fmt.Scan(&*respondgame)
		}
		switch *respondgame {
		case 1:
			ptile[0] = domino[help]
			count++
			help++
		case 2:
			ptile[1] = domino[help]
			count++
			help++
		case 3:
			ptile[2] = domino[help]
			count++
			help++
		case 4:
			ptile[3] = domino[help]
			count++
			help++
		}
		if count == 2 {
			*respondgame = 0
		}
	}
}

// FOR PLAYER

func namecheck(p *playerinfo, pidx *int) {
	// a procedure to check if a name inputted had already exist before
	var exist bool = false
	var confirm string
	var new bool
	var i int
	for i = 0; i < numplayers && exist == false; i++ {
		if player[i].name == p.name {
			exist = true
		}
	}

	// if the player wants to start anew
	if exist == true {
		fmt.Print("Your name had already exist. Would you like to import your previous stats ? (y/n)\n", "--> ")
		fmt.Scan(&confirm)
		for confirm != "y" && confirm != "n" {
			fmt.Print("Invalid input.\n", "The valid inputs are y and n.\n", "Try inputting again.\n", "--> ")
			fmt.Scan(&confirm)
		}
	}

	if confirm == "y" {
		new = false
	} else {
		new = true
	}

	if exist == true && new == false {
		*p = player[i-1]
		*pidx = i - 1
	} else if exist == true && new == true {
		p.score = 0
		p.game = 0
		*pidx = i - 1
	} else {
		if numplayers < N {
			player[numplayers].name = p.name
			*pidx = numplayers
			numplayers++
		} else {
			fmt.Print("Sorry, you can't register any more player.\n")
			fmt.Print("Please insert an already existing name. Note that names are case-sensitive.\n", "--> ")
			fmt.Scan(&p.name)
			namecheck(&*p, &*pidx)
		}
	}
}

func gameresult(ptile [4]domtile, dtile [4]domtile, p *playerinfo) {
	// a procedure to calculate the result of the game
	var pbal, dbal, ptotbal, dtotbal, psum, dsum int
	for i := 0; i < 4; i++ {
		psum += ptile[i].left + ptile[i].right
		dsum += dtile[i].left + dtile[i].right
		if ptile[i].balak == true {
			pbal++
			ptotbal += ptile[i].left
		}
		if dtile[i].balak == true {
			dbal++
			dtotbal += dtile[i].left
		}
	}
	if pbal >= 2 && dbal < 2 {
		// displaying the result and the reason why
		fmt.Print("You have ", pbal, " (2 or more) double tiles while the dealer has ", dbal, ".\n")
		fmt.Println("=> You won the round.")
		p.game++
		p.score++
	} else if pbal < 2 && dbal >= 2 {
		fmt.Print("You have ", pbal, " (less than 2) double tile while the dealer has ", dbal, ".\n")
		fmt.Println("=> You lost the round.")
		p.game++
	} else if pbal >= 2 && dbal >= 2 {
		if ptotbal > dtotbal {
			fmt.Print("Both you and the dealer have ", pbal, " (2 or more) double tiles.\n", "The sum of the dots on your double tiles is higher than the dealer's.\n")
			fmt.Println("=> You won the round.")
			p.game++
			p.score++
		} else if ptotbal < dtotbal {
			fmt.Print("Both you and the dealer have ", pbal, " (2 or more) double tiles.\n", "The sum of the dots on the dealer's double tiles is higher than yours.\n")
			fmt.Println("=> You lost the round.")
			p.game++
		} else {
			if psum > dsum {
				fmt.Print("Both you and the dealer have ", pbal, " (2 or more) double tiles.\n", "The sum of the dots on the double tiles is also a tie.\n", "The sum of the dots on all your tiles is higher than the dealer's.\n")
				fmt.Println("=> You won the round.")
				p.game++
				p.score++
			} else if psum < dsum {
				fmt.Print("Both you and the dealer have ", pbal, " (2 or more) double tiles.\n", "The sum of the dots on the double tiles is also a tie.\n", "The sum of the dots on all the dealer's tiles is higher than yours.\n")
				fmt.Println("=> You lost the round.")
				p.game++
			} else {
				fmt.Println("Both you and the dealer have 2 or more double tiles and the same sum of dots on the double tiles and the overall tiles.")
				fmt.Println("=> The round's a tie.")
				p.game++
			}
		}
	} else {
		if psum > dsum {
			fmt.Print("Both you and the dealer have less than 2 double tiles.\n", "The sum of dots on all your tiles is higher than the dealer's.\n")
			fmt.Println("=> You won the round.")
			p.game++
			p.score++
		} else if psum < dsum {
			fmt.Print("Both you and the dealer have less than 2 double tiles.\n", "The sum of dots on all the dealer's tiles is higher than yours.\n")
			fmt.Println("=> You lost the round.")
			p.game++
		} else {
			fmt.Print("Both you and the dealer have less than 2 double tiles.\n", "The sum of all dots on the tiles is also the same beetwen you and the dealer.\n")
			fmt.Println("=> The round's a tie.")
			p.game++
		}
	}
	printscore(*p)
	AnyKey("Press any key and then enter to continue... ")
	clear()
	fmt.Println(p.name)
	if p.score > 1 && p.game > 1 {
		fmt.Print("Score : ", p.score, " wins from ", p.game, " games.\n")
	} else if p.score > 1 && p.game <= 1 {
		fmt.Print("Score : ", p.score, " wins from ", p.game, " game.\n")
	} else if p.score <= 1 && p.game > 1 {
		fmt.Print("Score : ", p.score, " win from ", p.game, " games.\n")
	} else {
		fmt.Print("Score : ", p.score, " win from ", p.game, " game.\n")
	}
	fmt.Println("---------------------------------------------------------------------------------------------------NEXTGAME")
}

func showscoreboard() {
	// a procedure to display the scoreboard
	var temp playerinfo
	var max, low int
	fmt.Println("**SCOREBOARD**")
	// calculate the winning rate of every player
	for i := 0; i < numplayers; i++ {
		player[i].rate = (float32(player[i].score) / float32(player[i].game)) * 100
		if player[i].game == 0 {
			player[i].rate = 0
		}
	}
	// sorting the rankings using selection sort
	for low = 0; low < numplayers-1; low++ {
		max = low
		for i := low + 1; i < numplayers; i++ {
			if player[i].rate > player[max].rate {
				max = i
			}
		}
		temp = player[low]
		player[low] = player[max]
		player[max] = temp
	}

	for i := 0; i < numplayers; i++ {
		// displaying the rankings
		fmt.Print("#", i+1, " ", player[i].name, ", with win percentage of ", player[i].rate, "%\n")
	}
	if numplayers > 0 {
		fmt.Println("Note that this list will disappear after you close the game.")
	} else {
		fmt.Println("There is no player registered. Play a game to see the scoreboard.")
	}
	AnyKey("Press any key and then enter to go back... ")
}

func instructions() {
	// a procedure to display how the game works
	fmt.Println("**HOW TO PLAY**")
	fmt.Println("1. Dominos have 28 different kind of tiles with the dots on the left and right side ranging from 0 - 6. ")
	fmt.Println("2. Tiles that have the same number of dots on the left and the right side are called double tiles.")
	fmt.Println("3. Dominoes are randomized with the position of the domino facing down to form a Boneyard (randomized dominoes).")
	fmt.Println("4. Player can take four tiles from Boneyard with the position facing up.")
	fmt.Println("5. The opponent can also take four tiles from Boneyard with the condition facing down.")
	fmt.Println("6. Player can change tiles from Boneyard maximum twice.")
	fmt.Println("7. The winner is determined based on the following:")
	fmt.Println("	a. If the player has 2 or more double tiles and the opponent has less than 2, the player wins.")
	fmt.Println("	b. If the opponent also has 2 or more double tiles, the winner is determined by the sum of the dots on the double tiles.")
	fmt.Println("	c. If the sum of dots on the double tiles is a tie, the winner is determined by the sum of dots on all 4 tiles.")
	fmt.Println("	d. If the sum of dots on all 4 tiles is a tie, then the round is a tie.")
	fmt.Println("	e. If the player and the opponent have less than 2 double tiles, then use rule c and d to determine the winner.")
	fmt.Println(" ")
	fmt.Println("Pro Tip : Why not learn as you play ? It's more fun that way ;)")
	AnyKey("Press any key and then enter to go back... ")
}

func AnyKey(text string) {
	// a procedure to handle the "Press any key" part
	var any string
	fmt.Println(" ")
	fmt.Print(text)
	fmt.Scan(&any)
}

func clear() {
	// a procedure to clear the screen
	c := exec.Command("cmd", "/c", "cls")
	c.Stdout = os.Stdout
	c.Run()
}
