
export const calculateDuration = (startTime: string, endTime: string): string => {
    // Parse the start and end times
    const start = new Date(startTime);
    const end = new Date(endTime);

    // Calculate the difference in milliseconds
    const difference = end.getTime() - start.getTime();

    // Convert milliseconds to hours, minutes, and seconds
    const hours = Math.floor(difference / (1000 * 60 * 60));
    const minutes = Math.floor((difference % (1000 * 60 * 60)) / (1000 * 60));
    const seconds = Math.floor((difference % (1000 * 60)) / 1000);

    return printPrettyDuration(hours, minutes, seconds)
};
export const convertToSimpleDate = (isoDateTime: string): string => {
    const date = new Date(isoDateTime);

    const year = date.getFullYear();
    // Months in JavaScript are 0-indexed, so add 1 for the correct month number
    const month = (date.getMonth() + 1).toString().padStart(2, '0');
    const day = date.getDate().toString().padStart(2, '0');
    const hours = date.getHours().toString().padStart(2, '0');
    const minutes = date.getMinutes().toString().padStart(2, '0');

    return `${year}-${month}-${day} ${hours}:${minutes}`;
};

export const timeAgo = (dateTimeString: string): string => {
    // Parse the provided date string
    const pastDate = new Date(dateTimeString);
    const currentDate = new Date();

    // Calculate the difference in milliseconds
    const differenceInMilliseconds = currentDate.getTime() - pastDate.getTime();

    // Convert milliseconds to hours, minutes, and seconds
    const hoursAgo = Math.floor(differenceInMilliseconds / (1000 * 60 * 60));
    const minutesAgo = Math.floor((differenceInMilliseconds % (1000 * 60 * 60)) / (1000 * 60));
    const secondsAgo = Math.floor((differenceInMilliseconds % (1000 * 60)) / 1000);

    return `${printPrettyDuration(hoursAgo, minutesAgo, secondsAgo)} ago`;
};

const printPrettyDuration = (hours: number, minutes: number, seconds: number): string => {
    let durationParts = [];
    if (hours > 0) {
        durationParts.push(`${hours} hr${isPlural(hours)}`);
    }
    if (minutes > 0) {
        durationParts.push(`${minutes} min${isPlural(minutes)}`);
    }
    if (seconds > 0 || durationParts.length === 0) {
        durationParts.push(`${seconds} second${isPlural(seconds)}`);
    }
    return durationParts.join(", ");
}

const isPlural = (value: number) => value > 1 ? 's' : '';
