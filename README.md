## Intégration dans GraphShield

Ce module ALTCHA est autonome.

Endpoints disponibles :

- GET /health
- GET /api/altcha/challenge
- POST /api/altcha/verify

ALTCHA doit être déclenché lorsque le moteur décisionnel retourne :

CHALLENGE

### Frontend

Un exemple fonctionnel utilisant le widget ALTCHA officiel est disponible dans :

static/index.html

Le widget utilise :

GET /api/altcha/challenge

Après résolution, le payload généré doit être envoyé vers :

POST /api/altcha/verify

Format :

{
  "payload": "<ALTCHA_PAYLOAD>"
}

Réponse valide :

{
  "verified": true
}

### Configuration

Variable obligatoire :

ALTCHA_HMAC_SECRET

Exemple Docker :

docker run -d \
  --name altcha-server \
  -p 8080:8080 \
  -e ALTCHA_HMAC_SECRET="YOUR_SECRET" \
  altcha-server