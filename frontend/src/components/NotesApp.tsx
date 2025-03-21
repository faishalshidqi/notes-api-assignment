import NotesAppHeader from './NotesAppHeader.tsx'
import HomePage from '../pages/HomePage.tsx'
import {Route, Routes} from 'react-router-dom'
import DetailPageWrapper from '../pages/NoteDetailPage.tsx'
import AddNotePage from '../pages/AddNotePage.tsx'
import NotFoundPage from '../pages/NotFoundPage.tsx'
import {useState, useEffect, useMemo} from 'react'
import LoginPage from '../pages/LoginPage.tsx'
import RegisterPage from '../pages/RegisterPage.tsx'
import {getUserLogged, putAccessToken} from '../utils/data.ts'
import {UserAuthProvider} from '../contexts/UserAuthContext.ts'
import {LocaleProvider} from '../contexts/LocaleContext.ts'
import {ThemeProvider} from '../contexts/ThemeContext.ts'
import Loading from './Loading.tsx'

export default function NotesApp() {
    const [user, setUser] = useState<{id: string, name: string, username: string}>({id: '', name: '', username: ''})
    const [locale, setLocale] = useState<string>('en')
    const [theme, setTheme] = useState<string>('light')
    const [loading, setLoading] = useState<boolean>(true)
    const userAuthContextValue = useMemo(() => {
        return {
            user
        }
    }, [user])
    const localeContextValue = useMemo(() => {
        return {
            locale: localStorage.getItem('locale') || 'en',
            toggleLocale: () => {
                const newLocale = localStorage.getItem('locale') === 'en' ? 'id' : 'en'
                localStorage.setItem('locale', newLocale)
                setLocale(newLocale)
            }
        }
    }, [locale])
    const themeContextValue = useMemo(() => {
        return {
            theme: localStorage.getItem('theme') || 'light',
            toggleTheme: () => {
                const newTheme = localStorage.getItem('theme') === 'light' ? 'dark' : 'light'
                localStorage.setItem('theme', newTheme)
                setTheme(newTheme)
            }
        }
    }, [theme])
    async function onSuccessLogin({accessToken}: {accessToken: string}) {
        putAccessToken(accessToken)
        const {data} = await getUserLogged()
        setUser(data)
    }
    const onLogout = () => {
        setUser({id: '', name: '', username: ''})
        putAccessToken('')
    }

    useEffect(() => {
        getUserLogged().then(({error, data}) => {
            if (error) (setUser({id: '', name: '', username: ''}))
            if (!error) setUser(data)
            setLoading(false)
        })
    }, [])

    useEffect(() => {
        document.documentElement.setAttribute('data-theme', theme)
    }, [theme])

    const not_authed = (
        <ThemeProvider value={themeContextValue}>
            <LocaleProvider value={localeContextValue}>
                <UserAuthProvider value={userAuthContextValue.user}>
                    <div className='app-container'>
                        <NotesAppHeader onLogout={onLogout} />
                        <main>
                            <Routes>
                                <Route path='/' element={<LoginPage onSuccessLogin={onSuccessLogin}/>} />
                                <Route path='/register' element={<RegisterPage/>} />
                                <Route path='*' element={<NotFoundPage/>}/>
                            </Routes>
                        </main>
                    </div>
                </UserAuthProvider>
            </LocaleProvider>
        </ThemeProvider>
    )

    const authed = (
        <ThemeProvider value={themeContextValue}>
            <LocaleProvider value={localeContextValue}>
                <UserAuthProvider value={user}>
                    <div className='app-container'>
                        <NotesAppHeader onLogout={onLogout} />
                        <main>
                            <Routes>
                                <Route path='/' element={<HomePage/>} />
                                <Route path='/notes/app/:id' element={<DetailPageWrapper/>} />
                                <Route path='/notes/new' element={<AddNotePage/>}/>
                                <Route path='*' element={<NotFoundPage/>}/>
                            </Routes>
                        </main>
                    </div>
                </UserAuthProvider>
            </LocaleProvider>
        </ThemeProvider>
    )
    if (loading) {
        return <Loading />
    }

    return user.id === '' ? not_authed : authed
}