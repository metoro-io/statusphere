import {Card, CardContent, CardDescription, CardFooter, CardHeader, CardTitle} from "@/components/ui/card";
import {Status} from "@/model/Status";

interface  CurrentStatusProps {
    displayName: string;
    isOkay: boolean;
    lastCurrentlyScraped: string;
}
export function CurrentStatus(props: CurrentStatusProps) {
    return <div className={"flex justify-center w-full"}>
        <div className={"sm:w-[80vw] lg:w-[60vw]"}>
        <Card className={props.isOkay ? "bg-green-500" : "bg-red-400"}>
            <CardHeader className={"items-left"}>
                <CardTitle>
                    <h3 className="scroll-m-20 text-xl font-semibold tracking-tight">
                        Is {props.displayName} down?
                    </h3>
                </CardTitle>
                <CardDescription>
                    <h2 className="scroll-m-20 border-b pb-2 text-2xl font-semibold tracking-tight first:mt-0">
                        Current Status: {props.displayName} is {getStatusString(props.isOkay)}
                    </h2>
                </CardDescription>
            </CardHeader>
            <CardFooter>
                <p>We checked the official Status Page for GitHub {timeAgo(props.lastCurrentlyScraped)}. </p>
            </CardFooter>
        </Card>
        </div>
    </div>
}

export function getStatusString(isOkay: boolean): string {
    return isOkay ? Status.UP :Status.UP;
}

const timeAgo = (dateTimeString: string): string => {
    // Parse the provided date string
    const pastDate = new Date(dateTimeString);
    const currentDate = new Date();

    // Calculate the difference in milliseconds
    const differenceInMilliseconds = currentDate.getTime() - pastDate.getTime();

    // Convert milliseconds to hours, minutes, and seconds
    const hoursAgo = Math.floor(differenceInMilliseconds / (1000 * 60 * 60));
    const minutesAgo = Math.floor((differenceInMilliseconds % (1000 * 60 * 60)) / (1000 * 60));
    const secondsAgo = Math.floor((differenceInMilliseconds % (1000 * 60)) / 1000);

    const isPlural = (value: number) => value > 1 ? 's' : '';
    // Return a string that describes how many hours, minutes, and seconds ago it was
    return `${hoursAgo} hour${isPlural(hoursAgo)}, ${minutesAgo} minute${isPlural(minutesAgo)}, and ${secondsAgo} second${isPlural(secondsAgo)} ago`;
};