import NoteItem from './NoteItem.tsx'

export default function NotesList({notes, onDelete, onArchive}: {notes: {id: string, title: string, body: string, createdAt: string, archived: boolean}[], onArchive: (id: string) => void, onDelete: (id: string) => void}) {
    if (notes.length === 0) {
        return <section className='notes-list-empty'><p>There are no notes here!</p></section>
    }
    return (
        <section className='notes-list'>
            {
                notes.map((note) => (
                    <NoteItem {...note} key={note.id} onDelete={onDelete} onArchive={onArchive} />
                ))
            }
        </section>
    )
}