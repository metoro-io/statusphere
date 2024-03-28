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

    return (
        <>
            <div className="flex w-full items-center space-x-2 mt-4">
                <Popover open={open} onOpenChange={(a) => {
                    setOpen(a)
                }}>
                    <PopoverTrigger asChild>
                        <Button
                            variant="outline"
                            role="combobox"
                            aria-expanded={open}
                            className="w-full justify-between bg-white shadow-white text-gray-400 hover:bg-slate-100 hover:text-slate-900"
                        > {company === "" ? "Search company status" : company}
                            <div className={"flex w-full justify-end "}>
                                <CaretSortIcon className="ml-2 h-4 w-4 shrink-0 opacity-50"/>
                            </div>
                        </Button>
                    </PopoverTrigger>
                    <PopoverContent className="w-full p-0 popover-content-width-same-as-its-trigger bg-white">
                        <Command className={"w-full"}>
                            <CommandInput onValueChange={(a) => {
                                setPrefix(a)
                            }}
                                          placeholder="Type a company name..."/>
                            <CommandEmpty>No company found.</CommandEmpty>
                            <CommandGroup>
                                {companyList.map((details) => (
                                    <CommandItem
                                        className={"w-full text-slate-700 hover:bg-slate-100 hover:text-slate-900"}
                                        key={details.name}
                                        value={details.name}
                                        onSelect={(currentValue) => {
                                            setCompany(currentValue)
                                            setOpen(false)
                                        }}
                                    >
                                        {details.name}
                                    </CommandItem>
                                ))}
                            </CommandGroup>
                        </Command>
                    </PopoverContent>
                </Popover>
                <div>
                    <Button
                        className="w-full bg-white shadow-white border text-slate-700 hover:bg-slate-100 hover:text-slate-900"
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