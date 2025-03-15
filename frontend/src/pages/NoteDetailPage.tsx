import NoteDetail from '../components/NoteDetail.tsx'
import {Component} from 'react'
import {deleteNote, getNote} from '../utils/data.ts'
import {useNavigate, useParams} from 'react-router-dom'
import Loading from '../components/Loading.tsx'

export default function DetailPageWrapper() {
    const {id} = useParams()
    const navigate = useNavigate()
    async function onDeleteHandler() {
        if (id != null) {
            await deleteNote(id)
        }
        navigate('/')
    }
    return <DetailPage id={id ?? ''} onDelete={onDeleteHandler}/>
}

class DetailPage extends Component<{id: string, onDelete: (id: string) => void}, {note: {id: string, title: string, body: string, createdAt: string, archived: boolean}|undefined, loading: boolean}> {
    constructor(props: {id: string, onDelete: (id: string) => void, onArchive: (id: string) => void}) {
        super(props)
        this.state = {
            note: {id: '', title: '', body: '', createdAt: '', archived: false},
            loading: true
        }
        this.onDeleteHandler = this.onDeleteHandler.bind(this)
    }
    onDeleteHandler(id: string) {
        this.props.onDelete(id)
        this.setState(() => {
            return {
                loading: false
            }
        })
    }
    async componentDidMount() {
        const {data} = await getNote(this.props.id)
        this.setState(() => {
            return {
                note: data,
                loading: false
            }
        })
    }
    render() {
        if (this.state.loading) {
            return <Loading />
        }
        return (
            <section className='detail-page'>
                <NoteDetail onDelete={this.onDeleteHandler} id={this.state.note?.id ?? ''} title={this.state.note?.title ?? ''} body={this.state.note?.body ?? ''} createdAt={this.state.note?.createdAt ?? ''} archived={this.state.note?.archived ?? false} />
            </section>
        )
    }
}