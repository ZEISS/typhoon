'use client'

import {
  Popover,
  PopoverTrigger,
  PopoverContent
} from '@/components/ui/popover'
import {
  Command,
  CommandEmpty,
  CommandInput,
  CommandList,
  CommandGroup,
  CommandItem
} from '@/components/ui/command'
import {
  Dialog,
  DialogContent,
  DialogDescription,
  DialogFooter,
  DialogHeader,
  DialogTitle,
  DialogTrigger
} from '@/components/ui/dialog'
import {
  Form,
  FormControl,
  FormItem,
  FormLabel,
  FormDescription,
  FormMessage,
  FormField
} from '@/components/ui/form'
import { Input } from '@/components/ui/input'
import { CaretSortIcon } from '@radix-ui/react-icons'
import { Button } from '@/components/ui/button'
import { useState } from 'react'
import { useForm } from 'react-hook-form'
import { zodResolver } from '@hookform/resolvers/zod'
import { cn } from '@/lib/utils'
import { rhfCreateTeamAction } from '@/actions/teams'
import { Textarea } from '@/components/ui/textarea'
import {
  defaultValues,
  TeamCreateSchema,
  CreateTeamFornValues
} from '@/schemas/teams'

export function TeamSwitcher() {
  const [open, setOpen] = useState(false)
  const [showNewTeamDialog, setShowNewTeamDialog] = useState(false)

  const form = useForm<CreateTeamFornValues>({
    resolver: zodResolver(TeamCreateSchema),
    defaultValues,
    mode: 'onChange'
  })

  const handleSubmit = () => { }

  return (
    <Dialog open={showNewTeamDialog} onOpenChange={setShowNewTeamDialog}>
      <Popover open={open} onOpenChange={setOpen}>
        <PopoverTrigger asChild>
          <Button
            variant="outline"
            role="combobox"
            aria-expanded={open}
            aria-label="Select a team"
            className={cn('w-[200px] justify-between')}
          >
            No Team
            <CaretSortIcon className="ml-auto h-4 w-4 shrink-0 opacity-50" />
          </Button>
        </PopoverTrigger>
        <PopoverContent className="w-[200px] p-0">
          <Command>
            <CommandList>
              <CommandInput placeholder="Search team..." />
              <CommandEmpty>No team found.</CommandEmpty>
            </CommandList>
            <CommandList>
              <CommandGroup>
                <DialogTrigger asChild>
                  <CommandItem
                    onSelect={() => {
                      setOpen(false)
                      setShowNewTeamDialog(true)
                    }}
                  >
                    Create Team
                  </CommandItem>
                </DialogTrigger>
                <CommandItem>Manage Team</CommandItem>
              </CommandGroup>
            </CommandList>
          </Command>
        </PopoverContent>
      </Popover>
      <DialogContent>
        <DialogHeader>
          <DialogTitle>Create team</DialogTitle>
          <DialogDescription>
            Add a new team to manage products and customers.
          </DialogDescription>
        </DialogHeader>
        <Form {...form}>
          <form
            action={rhfCreateTeamAction}
            onSubmit={handleSubmit}
            className="space-y-8"
          >
            <FormField
              control={form.control}
              name="name"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="sr-only">
                    <h1>Name</h1>
                  </FormLabel>
                  <FormControl>
                    <Input placeholder="Name ..." {...field} />
                  </FormControl>
                  <FormDescription>Give it a great name.</FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="slug"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="sr-only">Slug</FormLabel>
                  <FormControl>
                    <Input placeholder="Slug ..." {...field} />
                  </FormControl>
                  <FormDescription>
                    {`This is the short name used for URLs (e.g.
                'solution-architects', 'order-service')`}
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="contactEmail"
              render={({ field }) => (
                <FormItem>
                  <FormLabel className="sr-only">
                    <h1>Contact email</h1>
                  </FormLabel>
                  <FormControl>
                    <Input placeholder="team@acme.com" {...field} />
                  </FormControl>
                  <FormDescription>
                    Add a shared inbox for you team (optional).
                  </FormDescription>
                  <FormMessage />
                </FormItem>
              )}
            />

            <FormField
              control={form.control}
              name="description"
              render={({ field }) => (
                <div className="grid w-full">
                  <FormItem>
                    <FormLabel className="sr-only">
                      <h1>Description</h1>
                    </FormLabel>
                    <FormControl>
                      <Textarea
                        {...field}
                        className="w-full"
                        placeholder="Add a description ..."
                      />
                    </FormControl>
                    <FormDescription>A desciption of your team</FormDescription>
                    <FormMessage />
                  </FormItem>
                </div>
              )}
            />

            <DialogFooter>
              <Button
                variant="outline"
                onClick={() => setShowNewTeamDialog(false)}
              >
                Cancel
              </Button>
              <Button
                type="submit"
                disabled={
                  form.formState.isSubmitting || !form.formState.isValid
                }
              >
                Continue
              </Button>
            </DialogFooter>
          </form>
        </Form>
      </DialogContent>
    </Dialog>
  )
}
