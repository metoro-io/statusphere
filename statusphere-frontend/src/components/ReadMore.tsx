import {useState} from "react";

interface ReadMoreProps {
    text: string | undefined | null;
}

const CHAR_LIMIT = 113
export const ReadMore = ({text}: ReadMoreProps) => {
    if (!text) {
        return <p>No description</p>
    }

    const [isExpanded, setIsExpanded] = useState(false)
    const subText = text.substring(0, CHAR_LIMIT)
    const shouldShowReadMore = text.length > CHAR_LIMIT

    return (
        <p>
            {shouldShowReadMore && !isExpanded ? subText + '...' : text}
            {!shouldShowReadMore && text}
            <span
                className='text-gray-400 underline ml-2'
                role="button"
                tabIndex={0}
                aria-expanded={isExpanded}
                onClick={() => setIsExpanded(!isExpanded)}
            >
                {shouldShowReadMore && (isExpanded ? 'show less' : 'read more')}
          </span>
        </p>
    )
}