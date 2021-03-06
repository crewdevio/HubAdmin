# HUB ADMIN

Hub Admin is a command-line tool managing your github credentials

## Installation

``` console
  
  go get github.com/crewdevio/HubAdmin
 
```

## How to use

### Open help
```console
hubadmin help
```
![image](https://user-images.githubusercontent.com/59743950/137951279-c7489d69-325f-414d-9d53-4ab865555432.png)

### Add your credentials

```console
hubadmin add <username> <access_token>
```
![image](https://user-images.githubusercontent.com/59743950/137822738-28bddf4a-f8f5-4629-8d13-d9ea9d144cac.png)

### List github accounts
```console 
hubadmin list
```
![image](https://user-images.githubusercontent.com/59743950/137949721-e224dd5f-a170-41dc-b89b-8fa1daa73228.png)

### Switch between your github accounts
``` console
  hubadmin use <username>
```
![image](https://user-images.githubusercontent.com/59743950/137822403-99c8fe17-4b4a-4e3a-91d5-6faaae550105.png)

### Delete your stored github account
```console
hubadmin delete <username>
```
![image](https://user-images.githubusercontent.com/59743950/137940254-23ad9e1c-0e75-4175-b6be-b9ba5e631db7.png)

## Support 
cli for the moment only works on windows
