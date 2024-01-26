'use client'

import { Dialog } from '@/components/ui/dialog'
import {
  Popover,
  PopoverTrigger,
  PopoverContent
} from '@/components/ui/popover'
import {
  Command,
  CommandEmpty,
  CommandInput,
  CommandList
} from '@/components/ui/command'
import { CaretSortIcon } from '@radix-ui/react-icons'
import { Button } from '@/components/ui/button'
import { useState } from 'react'
import { cn } from '@/lib/utils'

export function SystemSwitcher() {
  const [open, setOpen] = useState(false)

  return (
    <Dialog>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a system"
            className={cn('w-[200px] justify-between')}
          >
            Global System
            <CaretSortIcon className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command>
            <CommandList>
              <CommandInput placeholder="Search system..." />
              <CommandEmpty>No system found.</CommandEmpty>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
    </Dialog>
  )
}
