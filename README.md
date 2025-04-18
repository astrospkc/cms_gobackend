## One-Go

### Todos
- [X] Build schemas for User, project, blog, media,link,resume,videocontent,
- [X] Payment schemas- subscriptionPlan, usersubscription, Security/ api key schema- apikey
- [X] make routes for user(CRUD)
- [X] getuser, CRUD for projects, github links, links, media, usersubscription, apikey, blog category
- [ ] add aws s3 for file uploads
- [ ] read about security field to network request
- [ ] login page, signup , dashboard the frontend part.
- [ ] payment intergration 
- [ ] api key integration

###### Flow: Login with Email + Email Confirmation
- User enters email on frontend
- Backend generates a one-time token or link
- Token is emailed to the user
- User clicks the link (or enters the token)
- Backend verifies the token and logs them in