import {useState} from "react";

interface ReadMoreProps {
    text: string;
}

const CHAR_LIMIT = 113
export const ReadMore = ({text}: ReadMoreProps) => {
    const [isExpanded, setIsExpanded] = useState(false)
    const subText = text.substring(0, CHAR_LIMIT)

    return (
        <p>
            {subText}...
            <span
                className='text-gray-400 underline ml-2'
                role="button"
                tabIndex={0}
                aria-expanded={isExpanded}
                onClick={() => setIsExpanded(!isExpanded)}
            >
            {isExpanded ? 'show less' : 'read more'}
          </span>
        </p>
    )

}