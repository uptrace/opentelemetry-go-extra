#!/bin/bash
ENDPOINT='http://localhost:8080/graphql'

wget --body-data '{"query":"query AllUserFullNames {users {fullName}}","variables":{}}' \
  --method GET \
  $ENDPOINT
wget --body-data '{"query":"query SingleUserInfo($username: String!) {user(username: $username) {organization\nfullName\nusername\n}}","variables":{"username":"johnsmith"}}' \
  --method GET \
  $ENDPOINT
wget --body-data '{"query":"query UsersInOrganization($organization: String!){usersOfOrganization(organization: $organization) {username}}","variables":{"organization":"HR"}}' \
  --method GET \
  $ENDPOINT
wget --body-data '{"query":"mutation NewUser($user: UserInput!) {createUser(userInput: $user){username}}","variables":{"user":{"username":"joesmith","fullName":"Joe Smith","organization":"Marketing"}}}' \
  --method GET \
  $ENDPOINT
