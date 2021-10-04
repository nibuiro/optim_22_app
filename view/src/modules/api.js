// パスワードのハッシュ化
async function hashPassword(password) {
    const msgUint8 = new TextEncoder().encode(password); // パスワードをUint8Array(utf-8)としてエンコード
    const hashBuffer = await crypto.subtle.digest("SHA-256", msgUint8); // エンコードされたパスワードをハッシュ化
    const hashArray = Array.from(new Uint8Array(hashBuffer)); // バッファをbyte配列に変換
    const hashHex = hashArray
        .map(b => b.toString(16).padStart(2, "0"))
        .join(""); // byte配列を16進文字列に変換
    return hashHex;
}


// ユーザ登録API
async function register(component, user) {
    // パスワードのハッシュ化
    const hashHex = await hashPassword(user.password);
    // ログイン情報をサーバに送信し，レスポンスを得る
    const response = await fetch(`${process.env.API}/register`, {
        method: "POST",
        body: JSON.stringify({
            username: user.username,
            email: user.email,
            password: hashHex
        })
    });
    // 登録成功時
    if (response.status === 200) {
        const access_token = response.headers.get("Authorization");
        const refresh_token = response.headers.get("Refresh-Token");
        if (process.env.NODE_ENV === "development") {
            console.log("access_token:");
            console.log(access_token);
            console.log("refresh_token:");
            console.log(refresh_token);
        }
        // レスポンスのbodyをjsonに変換
        const data = await response.json();
        const user_id = data.user_id;
        if (process.env.NODE_ENV === "development") {
            console.log(`user_id: ${user_id}`);
        }
        // localStorageにユーザIDを保存
        localStorage.setItem("user_id", user_id);
        // localStorageにアクセストークンを保存
        localStorage.setItem("access_token", access_token);
        // cookieにリフレッシュトークンを保存（有効期限: 1ヶ月）
        component.$cookies.set("refresh_token", refresh_token, "1m");
        // 新規登録フォームを閉じる
        component.$emit("close");
        // ユーザ登録成功メッセージを表示する
        component.$emit("displayMessage");
    }
}


// ログインAPI
async function login(component, user) {
    // パスワードのハッシュ化
    const hashHex = await hashPassword(user.password);
    // ログイン情報をサーバに送信し，レスポンスを得る
    const response = await fetch(`${process.env.API}/auth`, {
        method: "POST",
        body: JSON.stringify({
            email: user.email,
            password: hashHex
        })
    });
    // ログイン成功時
    if (response.status === 200) {
        const access_token = response.headers.get("Authorization");
        const refresh_token = response.headers.get("Refresh-Token");
        if (process.env.NODE_ENV === "development") {
            console.log("access_token:");
            console.log(access_token);
            console.log("refresh_token:");
            console.log(refresh_token);
        }
        // レスポンスのbodyをjsonに変換
        const data = await response.json();
        const user_id = data.user_id;
        if (process.env.NODE_ENV === "development") {
            console.log(`user_id: ${user_id}`);
        }
        // localStorageにユーザIDを保存
        localStorage.setItem("user_id", user_id);
        // localStorageにアクセストークンを保存
        localStorage.setItem("access_token", access_token);
        // cookieにリフレッシュトークンを保存（有効期限: 1ヶ月）
        component.$cookies.set("refresh_token", refresh_token, "1m");
        // ログインフォームを閉じる
        component.$emit("close");
        // ページをリロード
        component.$router.push('/');
    } else {
        component.errorMessage = "ログインに失敗しました";
        component.invalid = true;
    }
}


// ユーザプロフィールの取得API
async function getProfile(user_id, access_token) {
    const response = await fetch(`${process.env.API}/user/${user_id}`, {
        method: 'GET',
        headers: {
            Authorization: access_token
        }
    });
    const profile = await response.json();
    if (process.env.NODE_ENV === "development") {
        console.log(`Profile of ${profile.username}:`);
        console.log(profile);
    }
    for (let i = 0; i < profile.submissions.length; i++) {
        const request = await getRequest(profile.submissions[i].request_id);
        profile.submissions[i].request = request;
        if (i === profile.submissions.length - 1) {
            return profile;
        }
    }
}


// ユーザプロフィールの編集API
async function editProfile(component, user, access_token) {
    // パスワードのハッシュ化
    const hashHex = await hashPassword(user.password);
    const profile = user;
    profile.password = hashHex;
    // プロフィール情報をサーバに送信し，レスポンスを得る
    const response = await fetch(`${process.env.API}/user/${profile.user_id}`, {
        method: "PUT",
        headers: {
            Authorization: access_token
        },
        body: JSON.stringify(profile)
    });
    // 登録成功時
    if (response.status === 200) {
        // 編集フォームを閉じる
        component.$emit("close");
        // ユーザ登録成功メッセージを表示する
        component.$emit("displayMessage");
    }
}


// リクエスト一覧取得API
async function getRequests() {
    const response = await fetch(`${process.env.API}/requests`);
    const requests = await response.json();
    if (process.env.NODE_ENV === "development") {
        console.log(`All Requests:`);
        console.log(requests.requests);
    }
    return requests.requests;
}


// 新規リクエスト投稿API
async function makeRequest(component, user_id, request, access_token) {
    const response = await fetch(`${process.env.API}/request`, {
        method: "POST",
        headers: {
            Authorization: access_token
        },
        body: JSON.stringify({
            requestname: request.title,
            client_id: user_id,
            content: request.detail
        })
    });
    // 登録成功時
    if (response.status === 200) {
        // 新規リクエストフォームを閉じる
        component.$emit("close");
        // ユーザ登録成功メッセージを表示する
        component.$emit("displayMessage");
    }
}


// リクエスト取得API
async function getRequest(request_id) {
    const response = await fetch(`${process.env.API}/request/${request_id}`);
    const request = await response.json();
    if (process.env.NODE_ENV === "development") {
        console.log(`Request #${request.request_id}:`);
        console.log(request);
    }
    return request;
}


// リクエスト編集API
async function editRequest(component, request, access_token) {
    // 提出物の情報をサーバに送信し，レスポンスを得る
    const response = await fetch(`${process.env.API}/request/${request.request_id}`, {
        method: "PUT",
        headers: {
            Authorization: access_token
        },
        body: JSON.stringify(request)
    });
    // 登録成功時
    if (response.status === 200) {
        // 編集フォームを閉じる
        component.$emit("close");
        // ユーザ登録成功メッセージを表示する
        component.$emit("displayMessage");
    }
}


// ディスカッション取得API
async function getComments(request_id) {
    const response = await fetch(`${process.env.API}/discussion/${request_id}`);
    const comments = await response.json();
    if (process.env.NODE_ENV === "development") {
        console.log(`Discussion of Request #${request_id}:`);
        console.log(comments.comments);
    }
    return comments.comments;
}


// コメント投稿API
async function addComment(component, comment, access_token) {
    const response = await fetch(`${process.env.API}/discussion/${comment.request_id}`, {
        method: "POST",
        headers: {
            Authorization: access_token
        },
        body: JSON.stringify(comment)
    });
    // 登録成功時
    if (response.status === 200) {
        // ユーザ登録成功メッセージを表示する
        return true;
    }
}


// サブミッション取得API
async function getsubmission(submission_id) {
    const response = await fetch(`${process.env.API}/submission/${submission_id}`);
    let submission = await response.json();
    if (process.env.NODE_ENV === "development") {
        console.log(`Submission #${submission.submission_id}:`);
        console.log(submission);
    }
    const request = await getRequest(submission.request_id);
    submission.request = request;
    return submission;
}


// サブミッション編集API
async function editProfile(component, submission, access_token) {
    // 提出物の情報をサーバに送信し，レスポンスを得る
    const response = await fetch(`${process.env.API}/submission/${submission.submission_id}`, {
        method: "PUT",
        headers: {
            Authorization: access_token
        },
        body: JSON.stringify(submission)
    });
    // 登録成功時
    if (response.status === 200) {
        // 編集フォームを閉じる
        component.$emit("close");
        // ユーザ登録成功メッセージを表示する
        component.$emit("displayMessage");
    }
}


export {
    register,
    login,
    getProfile,
    editProfile,
    getRequests,
    editRequest,
    makeRequest,
    getRequest,
    getComments,
    addComment,
    getsubmission,
    editSubmission
}