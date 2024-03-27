"use client"

import React, {useEffect, useState} from "react";
import {Button} from "@/components/ui/button";
import {Popover, PopoverContent, PopoverTrigger} from "@/components/ui/popover";
import {Command, CommandEmpty, CommandGroup, CommandInput, CommandItem} from "@/components/ui/command";
import {CaretSortIcon, CheckIcon} from "@radix-ui/react-icons";
import axios from "@/utils/axios";
import {cn} from "@/components/ui/lib/utils";
import {StatusPage} from "@/model/StatusPage";
import {useRouter} from "next/navigation";



export function Search() {
    const [company, setCompany] = useState<string>("");
    const [prefix, setPrefix] = useState<string>("");
    const [companyList, setCompanyList] = useState<StatusPage[]>([]);
    const [open, setOpen] = useState(false);
    const router = useRouter();

    useEffect(() => {
        const fetchData = async () => {
            try {
                const response = await axios.get(
                    '/api/v1/statusPages/search' + ((prefix.length === 0) ? "" :
                        ('?query=' + encodeURIComponent(prefix)))
                );
                if (response != undefined && response.data != undefined && response.data.statusPages != undefined) {
                    const companyList: StatusPage[] = response.data.statusPages
                    if (companyList.length != 0) {
                        setCompanyList(companyList.slice(0, 20));
                    }
                }

            } catch (err) {
                console.log(err);
            }
        };
        fetchData();
    }, [prefix]);

    const handleClick = () => {
    }

    return (
        <>
            <div className="flex w-full items-center space-x-2">
                    <Popover open={open} onOpenChange={(a) => {
                        setOpen(a)
                    }}>
                        <PopoverTrigger asChild>
                            <Button
                                variant="outline"
                                role="combobox"
                                aria-expanded={open}
                                className="w-full justify-between"
                            >
                                {company}

                                <div className={"flex w-full justify-end"}>
                                    <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50"/>
                                </div>
                            </Button>
                        </PopoverTrigger>
                        <PopoverContent className="w-full p-0 popover-content-width-same-as-its-trigger">
                            <Command className={"w-full"}>
                                <CommandInput onValueChange={(a) => {
                                    setPrefix(a)
                                }} placeholder="Search company..."
                                              className="h-9"/>
                                <CommandEmpty>No company found.</CommandEmpty>
                                <CommandGroup>
                                    {companyList.map((details) => (
                                        <CommandItem
                                            className={"w-full"}
                                            key={details.Name}
                                            value={details.Name}
                                            onSelect={(currentValue) => {
                                                setCompany(currentValue)
                                                setOpen(false)
                                            }}
                                        >
                                            {details.Name}
                                            <CheckIcon
                                                className={cn(
                                                    "ml-auto h-4 w-4",
                                                    company === details.Name ? "opacity-100" : "opacity-0"
                                                )}
                                            />
                                        </CommandItem>
                                    ))}
                                </CommandGroup>
                            </Command>
                        </PopoverContent>
                    </Popover>
                <div>
                    <Button
                        className="w-full"
                        onClick={() => {
                            router.push('/status/' + company)
                        }}
                        disabled={company === ""}
                    >
                        Search
                    </Button>
                </div>
            </div>
        </>
    );
}