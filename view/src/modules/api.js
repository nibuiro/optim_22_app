// ログインAPI
async function login(component, email, password) {
    const msgUint8 = new TextEncoder().encode(password); // パスワードをUint8Array(utf-8)としてエンコード
    const hashBuffer = await crypto.subtle.digest("SHA-256", msgUint8); // エンコードされたパスワードをハッシュ化
    const hashArray = Array.from(new Uint8Array(hashBuffer)); // バッファをbyte配列に変換
    const hashHex = hashArray
        .map(b => b.toString(16).padStart(2, "0"))
        .join(""); // byte配列を16進文字列に変換
    // ログイン情報をサーバに送信し，レスポンスを得る
    const response = await fetch(`${process.env.API}/auth`, {
        method: "POST",
        body: JSON.stringify({
            email: email,
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
        component.$router.go('/');
    } else {
        component.errorMessage = "ログインに失敗しました";
        component.invalid = true;
    }
}


// リクエスト一覧取得API
async function getRequests() {
    const response = await fetch(`${process.env.API}/requests`);
    const requests = await response.json();
    return requests.requests;
}

export {
    login,
    getRequests
}