const BASE_URL = "http://localhost:5000"

function getAccessToken() {
    return localStorage.getItem('accessToken')
}

function putAccessToken(accessToken: string) {
    return localStorage.setItem('accessToken', accessToken)
}

async function fetchWithToken(url: string, options: {headers?: object, method: 'GET'|'POST'|'PUT'|'DELETE', body?: string} = {headers: {}, method: 'GET'}) {
    return fetch(url, {
        ...options,
        headers: {
            ...options.headers,
            Authorization: `Bearer ${getAccessToken()}`,
        },
    })
}

async function login({ username, password }: { username: string, password: string }) {
    const response = await fetch(`${BASE_URL}/authentications`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ username, password }),
    })

    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        alert(responseJson.message)
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

async function register({ fullname, username, password }: {fullname: string, username: string, password: string }) {
    const response = await fetch(`${BASE_URL}/users`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ fullname, username, password }),
    })

    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        alert(responseJson.message)
        return { error: true }
    }

    return { error: false }
}

async function getUserLogged() {
    const response = await fetchWithToken(`${BASE_URL}/users`)
    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

async function addNote({ title, body }: { title: string, body: string }) {
    const response = await fetchWithToken(`${BASE_URL}/notes`, {
        method: 'POST',
        headers: {
            'Content-Type': 'application/json',
        },
        body: JSON.stringify({ title, body }),
    })

    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

async function getActiveNotes() {
    const response = await fetchWithToken(`${BASE_URL}/notes`)
    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

async function getNote(id: string) {
    const response = await fetchWithToken(`${BASE_URL}/notes/${id}`)
    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

async function archiveNote(id: string) {
    const response = await fetchWithToken(`${BASE_URL}/notes/${id}/archive`, {
        method: 'POST',
    })

    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

async function deleteNote(id: string) {
    const response = await fetchWithToken(`${BASE_URL}/notes/${id}`, {
        method: 'DELETE',
    })

    const responseJson = await response.json()

    if (responseJson.status !== 'success') {
        return { error: true, data: null }
    }

    return { error: false, data: responseJson.data }
}

const showFormattedDate = (date: string, localeString: string) => {
    return new Date(date).toLocaleDateString(localeString, {
        weekday: 'long',
        year: 'numeric',
        month: 'long',
        day: 'numeric',
    })
}

export {
    getAccessToken,
    putAccessToken,
    login,
    register,
    getUserLogged,
    getActiveNotes,
    deleteNote,
    getNote,
    archiveNote,
    addNote,
    showFormattedDate,
}
