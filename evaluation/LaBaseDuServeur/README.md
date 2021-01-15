1 La base du serveur

	. Créez un nouveau projet go.

	. Définissez une fonction main, qui sera la base de l’exécution de votre programme.

	. Définissez un type Task, qui est une struct avec deux champs:
		• ”Description”, de type string
		• ”Done”, de type bool

	. Définissez une variable globale ”task” qui est une slice de Task.

	. Définissez trois HandleFunc (https://golang.org/pkg/net/http/#HandleFunc) et donnez-leur respectivement:
		• Un path ”/”, et une fonction ”list”
		• Un path ”/done”, et une fonction ”done”
		• Un path ”/add”, et une fonction ”add”

	. Appelez, dans le main, la fonction ListenAndServe (https://golang.org/pkg/net/http/#ListenAndServe) qui vous permettra de lancer votre serveur HTTP avec les handlers que vous venez de définir.

