import {MdAdd} from 'react-icons/md'
import {Link} from 'react-router-dom'

export default function AddNoteButton() {
    return (
        <Link to='/notes/new'>
            <button className='action' type='button' title='Add Note'>
                <MdAdd/>
            </button>
        </Link>
    )
}