async function getRequests() {
    const response = await fetch(`${process.env.API}/requests`);
    const requests = await response.json();
    return requests.requests;
}

export {
    getRequests
}