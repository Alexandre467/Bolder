package main

import (
	f "fmt"
	"os"
)


////////////////////
/* Mes fonctions */
///////////////////


/* Fonction main */
func main() {
	myCMD()
}

/* Fonction qui contient tous les arguments de la commande bolder */
func myCMD() {
	lenghtarg := os.Args[1:]
	if len(lenghtarg) >= 1 {
		arg1 := os.Args[1]
		switch arg1 {
		case "-h" , "help":
			helpCMD()
		case "g", "get":
			getCMD()
		default:
			f.Println("Erreur de Syntaxe")
		}
	} 
}


/* Fonction get */
func getCMD() {
	lenghtarg := os.Args[1:]
	if len(lenghtarg) >= 2 {
		arg2 := os.Args[2]
		switch arg2 {
		case "ns", "namespaces":
			nsCMD()
		case "p", "pods":
			podCMD()
		case "n", "nodes":
			nodeCMD()
		default:
			f.Println("Erreur de Syntaxe")
		}
	}
}

/* Fonction help */
func helpCMD()  {
	myhelp :=  `Aide de Bolder
	-h, help : Liste des commandes
	g, get : Afficher les informations dune ressource
		- Namespaces, ns :
			Commande : bolder get namespaces
		- Pods :
			Commande : bolder get pods
		- Nodes, n :
			Commande : bolder get nodes`
			
	f.Println(myhelp)
}

/* Fonction qui liste les Namespaces */
func nsCMD() {

}

/* Fonction qui liste les pods */
func podCMD() {

}

/* Fonction qui liste les nodes */
func nodeCMD() {

}