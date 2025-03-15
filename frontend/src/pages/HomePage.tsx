import NotesList from '../components/NotesList.tsx'
import {Component} from 'react'
import AddNoteButton from '../components/AddNoteButton.tsx'
import {archiveNote, deleteNote, getActiveNotes} from '../utils/data.ts'
import SearchBar from '../components/SearchBar.tsx'
import {LocaleConsumer} from '../contexts/LocaleContext.ts'
import Loading from '../components/Loading.tsx'

export default class HomePage extends Component<unknown, { notes: { id: string, title: string, body: string, createdAt: string, archived: boolean }[], keyword: string, loading: boolean }> {
    constructor(props: unknown) {
        super(props)
        this.state = {
            notes: [],
            keyword: '',
            loading: true,
        }
        this.onDeleteHandler = this.onDeleteHandler.bind(this)
        this.onArchiveHandler = this.onArchiveHandler.bind(this)
        this.onSearchHandler = this.onSearchHandler.bind(this)
    }
    onSearchHandler({keyword}: { keyword: string }) {
        this.setState(() => ({keyword: keyword}))
    }
    async onDeleteHandler(id: string){
        await deleteNote(id)
        const {data} = await getActiveNotes()
        this.setState(() => {
            return {
                notes: data,
                loading: false
            }
        })

    }
    async onArchiveHandler(id: string){
        await archiveNote(id)
        const {data} = await getActiveNotes()
        this.setState(() => {
            return {
                notes: data,
                loading: false
            }
        })
    }
    async componentDidMount() {
        const {data} = await getActiveNotes()
        this.setState(() => {
            return {
                notes: data,
                loading: false
            }
        })
    }
    render() {
        if (this.state.loading) {
            return <Loading/>
        }
        const notes = this.state.notes.filter(note => note.title.toLowerCase().includes(this.state.keyword.toLowerCase()))
        return (
            <LocaleConsumer>
                {
                    ({locale}) => {
                        return (
                            <section className='homepage'>
                                <h2>{locale === 'en' ? 'Active Notes' : 'Catatan Aktif'}</h2>
                                <SearchBar searchNotes={this.onSearchHandler}/>
                                <NotesList notes={notes} onArchive={this.onArchiveHandler} onDelete={this.onDeleteHandler} />
                                <div className='homepage__action'>
                                    <AddNoteButton/>
                                </div>
                            </section>
                        )
                    }
                }
            </LocaleConsumer>
        )
    }
}