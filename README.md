# ESTIAM5_GO

Titre "Agent de voyage"
***

>On souhaite crée une application de agent de voyage.
>Un fois connecté sur l'application, il y aura possibiliter de choisir la destionation que vous suohaitez visiter depuis liste des destinations proposés.
>Après le chois de destination vous accèderez vers la page des hôtels proposés pour cette destination. 
>Vous pouvez choisir l'hôtel depuis list de hôtels proposé pour votre destination et ensuit vous pouvez reserver.


>Page connection/authentication: Utilisateur se connecte ou s'inscrit 

>Page accueil : Liste pays a visité proposés par agent de voyage, possibilité de choisir un pays cela redirige vers page Hôtels
 
>Page hôtels on a  nom de hôtel , adresse, description.




## EndPoint 
***


* GET /api/countrys

* GET /api/countrys/:id

* POST /api/countrys

---

* GET /api/hotels

* GET /api/hotel/:id

* GET /api/users/hotel/:iduser

- - -

* POST /api/user

* DELET /api/user/:id

* GET /api/users/:id

* GET /api/users

* POST /api/login

***

* GET /api/reservations/:id

* GET /api/reservations/hotels/:idUser

* POST /api/reservations/:idUser/:idHotel

* DELETE /api/reservations/:id


## Authors
***
Tatevik Havakemian
