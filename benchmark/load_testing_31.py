

#  1. POST   /api/user 
#  2. PUT    /api/user/:userID 
#  3. POST   /api/request 
#  4. PUT    /api/request/:request_id 
#  5. GET    /api/discussion/:requestID
#  6. GET    /api/user/:userID 
#  7. POST   /api/discussion/:requestID
#  8. GET    /api/submission/:submission_id
#  9. POST   /api/winner/:request_id 
#  
#  1. POST   /api/user 
#  2. GET    /api/requests
#  3. GET    /api/request/:request_id 
#  4. POST   /api/request/:request_id 
#  5. POST   /api/submission/:request_id
#  6. PUT    /api/submission/:submission_id
#  7. GET    /api/discussion/:requestID
#  8. POST   /api/discussion/:requestID

import requests
import uuid
import jwt
import json
from random import choice
from time import sleep
import mysql.connector




def ClientAgent():
  name = str(uuid.uuid4())
  response = Registration(name)
  access_token = get_access_token(response)
  user_id = get_user_id(access_token)
  EditProfile(name, user_id, access_token)
  title = str(uuid.uuid4())
  SubmitRequest(user_id, access_token, title)
  response = ViewRequests()
  request_id = get_request_id(response, title)
  EditRequest(access_token, request_id, title)
  ViewDiscussion(request_id)
  PostComment(access_token, request_id, user_id)
  yield request_id
  response = ViewRequest(request_id)
  engineer_id = choice_winner(response)
  submission_id = choice_submission_id(response)
  response = ViewWork(submission_id)
  ViewProfile(engineer_id)
  Like(access_token, request_id, engineer_id, user_id)
  yield


def EngineerAgent(request_id):
  name = str(uuid.uuid4())
  response = Registration(name)
  access_token = get_access_token(response)
  user_id = get_user_id(access_token)
  response = ViewRequests()
  request_id = request_id#choice_request_id(response)
  response = ViewRequest(request_id)
  response = JoinToRequest(access_token, request_id, user_id)
  response = SubmitWork(access_token, request_id, user_id)
  response = ViewRequest(request_id)
  submission_id = get_submission_id(response, user_id)
  response = ResubmitWork(access_token, submission_id, user_id)
  response = ViewDiscussion(request_id)
  PostComment(access_token, request_id, user_id)
  return


def Registration(name):
    url = "http://localhost:8080/api/user"
    json = {
        "username": name,
        "email": name+"@test.test",
        "password": "7f83b1657ff1fc53b992dc18148a1d6fffffd4b1fa3d677284addd20126d9069"
    }
    response = requests.post(url, json=json)
    return response


def EditProfile(name, user_id, access_token):
    url = f"http://localhost:8080/api/user/{user_id}"
    json = {
        "user_id":user_id,
        "username":"user1ss",
        "email":f"{name}@test.test",
        "password":"9f86d081884c7d659a2feaa0c55ad015a3bf4f1b2b0b822cd15d6c15b0f00a08",
        "icon":"data:image/png;base64,iVBORw0KGgoAAAANSUhEUgAAAAEAAAABCAIAAACQd1PeAAAAAXNSR0IArs4c6QAAAARnQU1BAACxjwv8YQUAAAAJcEhZcwAADsMAAA7DAcdvqGQAAAAMSURBVBhXY/j//z8ABf4C/qc1gYQAAAAASUVORK5CYII=",
        "comment":"C#",
        "sns": """{"github": "test","twitter": "test","facebook":"test"}"""
    }
    headers = {"Authorization": access_token}
    response = requests.put(url, json=json, headers=headers)
    return response


def ViewRequests():
    url = "http://localhost:8080/api/requests/"
    response = requests.get(url)
    return response


def get_request_id(response, title):
    body = json.loads(response.text)
    for row in body['requests']:
        if title == row['requestname']:
            return row['request_id']


def choice_request_id(response):
    body = json.loads(response.text)
    request = choice(body['requests'])
    return request['request_id']



def SubmitRequest(user_id, access_token, title):
    url = f"http://localhost:8080/api/request"
    json = {
        "client_id":str(user_id),
        "requestname":title,
        "content":"リクエスト0の詳細です。"
    }
    headers = {"Authorization": access_token}
    response = requests.post(url, json=json, headers=headers)
    return response
    

def EditRequest(access_token, request_id, title):
    url = f"http://localhost:8080/api/request/{request_id}"
    json = {
        "request_id":str(request_id),
        "requestname":title,
        "content":"xxxxxxx"
    }
    headers = {"Authorization": access_token}
    response = requests.put(url, json=json, headers=headers)
    return response


def ViewDiscussion(request_id):
    url = f"http://localhost:8080/api/discussion/{request_id}"
    response = requests.get(url)
    return response


def PostComment(access_token, request_id, user_id):
    url = f"http://localhost:8080/api/discussion/{request_id}"
    json = {
        "request_id":request_id,
        "user_id":user_id,
        "reply_id":0,
        "title":"コメント",
        "text":"リクエスト0に対するコメントです。",
        "attachment":"https://www.exampe.com"
    }
    headers = {"Authorization": access_token}
    response = requests.post(url, json=json, headers=headers)
    return response


def ViewProfile(user_id):
    url = f"http://localhost:8080/api/user/{user_id}"
    response = requests.get(url)
    return response



def ViewRequest(request_id):
    url = f"http://localhost:8080/api/request/{request_id}"
    response = requests.get(url)
    return response


def choice_submission_id(response):
    body = json.loads(response.text)
    submission = choice(body['request']['submissions'])
    return submission['submission_id']


def get_submission_id(response, user_id):
    body = json.loads(response.text)
    for submission in body['request']['submissions']:

        if user_id == submission['engineer']['user_id']:
            return submission['submission_id']


def choice_winner(response):
    body = json.loads(response.text)
    #print(body)
    submission = choice(body['request']['submissions'])
    return submission['engineer']['user_id']


def get_access_token(response):
    access_token = response.headers["Authorization"]
    return access_token


def get_user_id(access_token):
    claims = jwt.decode(
        access_token, 
        options={"verify_signature": False}, 
        algorithms=['HS256']
    )
    user_id = claims["userid"]
    return user_id


def ViewWork(submission_id):
    url = f"http://localhost:8080/api/submission/{submission_id}"
    response = requests.get(url)
    return response


def Like(access_token, request_id, engineer_id, client_id):
    url = f"http://localhost:8080/api/winner/{request_id}"
    json = {
        "request_id":str(request_id),
        "client_id":str(client_id),
        "engineer_id":str(engineer_id)
    }
    headers = {"Authorization": access_token}
    response = requests.post(url, json=json, headers=headers)
    return response


def JoinToRequest(access_token, request_id, user_id):
    url = f"http://localhost:8080/api/request/{request_id}"
    json = {
        "request_id":str(request_id),
        "engineer_id":str(user_id)
    }
    #print(json)
    headers = {"Authorization": access_token}
    response = requests.post(url, json=json, headers=headers)
    return response




def SubmitWork(access_token, request_id, user_id):
    url = f"http://localhost:8080/api/submission/{request_id}"
    json = {
        "request_id":str(request_id),
        "engineer_id":str(user_id),
        "url": "https://www.example.com/",
        "content": "エンジニア1の提出物です。"
    }
    headers = {"Authorization": access_token}
    response = requests.post(url, json=json, headers=headers)
    return response


def ResubmitWork(access_token, submission_id, user_id):
    url = f"http://localhost:8080/api/submission/{submission_id}"
    json = {
        "submission_id":str(submission_id),
        "engineer_id":str(user_id),
        "url": "https://www.example.com/xxxxxxx",
        "content": "エンジニア1の提出物です。編集済みです。"
    }
    headers = {"Authorization": access_token}
    response = requests.put(url, json=json, headers=headers)
    return response






if __name__ == '__main__':
    cnx = mysql.connector.connect(
        user='root', 
        password='rootpass',
        host='localhost',
        port=13306,
        database='optim_dev'
    )
    
    cur = cnx.cursor()
    for _ in range(10):

        cur.execute("""set foreign_key_checks = 0;""")
        cnx.commit()
        cur.execute("""truncate table requests;""")
        cnx.commit()
        cur.execute("""truncate table clients;""")
        cnx.commit()
        cur.execute("""truncate table engineer_requests;""")
        cnx.commit()
        cur.execute("""truncate table comments;""")
        cnx.commit()
        cur.execute("""truncate table engineers;""")
        cnx.commit()
        cur.execute("""truncate table profiles;""")
        cnx.commit()
        cur.execute("""truncate table submissions;""")
        cnx.commit()
        cur.execute("""truncate table users;""")
        cnx.commit()
        cur.execute("""truncate table winners;""")
        cnx.commit()
        cur.execute("""set foreign_key_checks = 1;""")
        cnx.commit()

        client_agent = ClientAgent()
        request_id = next(client_agent)
        EngineerAgent(request_id)
        next(client_agent)

    cnx.close()