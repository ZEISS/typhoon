package accounts

import (
	"context"

	"github.com/zeiss/typhoon/internal/models"
	"github.com/zeiss/typhoon/internal/web/components"
	"github.com/zeiss/typhoon/internal/web/ports"

	htmx "github.com/zeiss/fiber-htmx"
	"github.com/zeiss/fiber-htmx/components/alpine"
	"github.com/zeiss/fiber-htmx/components/buttons"
	"github.com/zeiss/fiber-htmx/components/cards"
	"github.com/zeiss/fiber-htmx/components/dropdowns"
	"github.com/zeiss/fiber-htmx/components/forms"
	"github.com/zeiss/fiber-htmx/components/icons"
	"github.com/zeiss/fiber-htmx/components/loading"
	"github.com/zeiss/fiber-htmx/components/tables"
	"github.com/zeiss/fiber-htmx/components/tailwind"
	"github.com/zeiss/pkg/conv"
)

// NewAccountControllerImpl ...
type NewAccountControllerImpl struct {
	Results tables.Results[models.Operator]
	Teams   tables.Results[models.Team]
	store   ports.Datastore
	htmx.DefaultController
}

// NewAccountController ...
func NewAccountController(store ports.Datastore) *NewAccountControllerImpl {
	return &NewAccountControllerImpl{store: store}
}

// Prepare ...
func (l *NewAccountControllerImpl) Prepare() error {
	return l.store.ReadTx(l.Context(), func(ctx context.Context, tx ports.ReadTx) error {
		err := tx.ListOperators(ctx, &l.Results)
		if err != nil {
			return err
		}

		return tx.ListTeams(ctx, &l.Teams)
	})
}

// Get ...
func (l *NewAccountControllerImpl) Get() error {
	return l.Render(
		components.DefaultLayout(
			components.DefaultLayoutProps{
				Title: "New Account",
				Path:  l.Path(),
			},
			func() htmx.Node {
				return htmx.FormElement(
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Properties"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								dropdowns.Dropdown(
									dropdowns.DropdownProps{},
									alpine.XData(`{
            filter: '',
            show: false,
            selected: null,
            focusedOptionIndex: null,
            options: null,
            close() { 
              this.show = false;
              this.filter = this.selectedName();
              this.focusedOptionIndex = this.selected ? this.focusedOptionIndex : null;
            },
            open() { 
              this.show = true; 
              this.filter = '';
            },
            toggle() { 
              if (this.show) {
                this.close();
              }
              else {
                this.open()
              }
            },
            isOpen() { return this.show === true },
            selectedName() { return this.selected ? this.selected.name.first + ' ' + this.selected.name.last : this.filter; },
            classOption(id, index) {
              const isSelected = this.selected ? (id == this.selected.login.uuid) : false;
              const isFocused = (index == this.focusedOptionIndex);
              return {
                'cursor-pointer w-full border-gray-100 border-b hover:bg-blue-50': true,
                'bg-blue-100': isSelected,
                'bg-blue-50': isFocused
              };
            },
            fetchOptions() {
              fetch('https://randomuser.me/api/?results=5')
                .then(response => response.json())
                .then(data => this.options = data);
            },
            filteredOptions() {
              return this.options
                ? this.options.results.filter(option => {
                    return (option.name.first.toLowerCase().indexOf(this.filter) > -1) 
                      || (option.name.last.toLowerCase().indexOf(this.filter) > -1)
                      || (option.email.toLowerCase().indexOf(this.filter) > -1)
                })
               : {}
            },
            onOptionClick(index) {
              this.focusedOptionIndex = index;
              this.selectOption();
            },
            selectOption() {
              this.focusedOptionIndex = this.focusedOptionIndex ?? 0;
              const selected = this.filteredOptions()[this.focusedOptionIndex]
              if (this.selected && this.selected.login.uuid == selected.login.uuid) {
                this.filter = '';
                this.selected = null;
              }
              else {
                this.selected = selected;
                this.filter = this.selectedName();
              }
              this.close();
            },
            focusPrevOption() {
              if (!this.isOpen()) {
                return;
              }
              const optionsNum = Object.keys(this.filteredOptions()).length - 1;
              if (this.focusedOptionIndex > 0 && this.focusedOptionIndex <= optionsNum) {
                this.focusedOptionIndex--;
              }
              else if (this.focusedOptionIndex == 0) {
                this.focusedOptionIndex = optionsNum;
              }
            },
            focusNextOption() {
              const optionsNum = Object.keys(this.filteredOptions()).length - 1;
              if (!this.isOpen()) {
                this.open();
              }
              if (this.focusedOptionIndex == null || this.focusedOptionIndex == optionsNum) {
                this.focusedOptionIndex = 0;
              }
              else if (this.focusedOptionIndex >= 0 && this.focusedOptionIndex < optionsNum) {
                this.focusedOptionIndex++;
              }
            }
        }`),
									alpine.XInit("fetchOptions()"),
									forms.TextInputBordered(
										forms.TextInputProps{
											Placeholder: "Please select a team",
										},
										alpine.XModel("filter"),
										loading.Spinner(
											loading.SpinnerProps{},
										),
									),
									dropdowns.DropdownMenuItems(
										dropdowns.DropdownMenuItemsProps{
											ClassNames: htmx.ClassNames{
												tailwind.WFull: true,
											},
										},

										htmx.Template(
											alpine.XFor("(option, index) in filteredOptions()"),
											dropdowns.DropdownMenuItem(
												dropdowns.DropdownMenuItemProps{},
												htmx.Div(
													alpine.XOn("click", "onOptionClick(index)"),
													alpine.XText("`${option.name.first}`"),
												),
											),
										),
									),
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Team"),
									),
								),
								forms.SelectBordered(
									forms.SelectProps{},
									htmx.Required(),
									htmx.HxValidate(true),
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select a team"),
									),
									htmx.Name("team_id"),
									htmx.Group(
										htmx.ForEach(l.Teams.GetRows(), func(operator *models.Team, idx int) htmx.Node {
											return forms.Option(
												forms.OptionProps{
													Value: operator.ID.String(),
												},
												htmx.Text(operator.Name),
											)
										})...,
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{},
								},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Operator"),
									),
								),
								forms.SelectBordered(
									forms.SelectProps{},
									htmx.HxGet("/accounts/partials/operator-skgs"),
									htmx.HxTarget("#operator-skgs"),
									htmx.HxSwap("outerHTML"),
									htmx.HxValidate(true),
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select an operator"),
									),
									htmx.Name("operator_id"),
									htmx.Group(
										htmx.ForEach(l.Results.GetRows(), func(operator *models.Operator, idx int) htmx.Node {
											return forms.Option(
												forms.OptionProps{
													Value: operator.ID.String(),
												},
												htmx.Text(operator.Name),
											)
										})...,
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Signing Key Group"),
									),
								),
								forms.SelectBordered(
									forms.SelectProps{},
									forms.Option(
										forms.OptionProps{
											Selected: true,
											Disabled: true,
										},
										htmx.Text("Select an signing key group"),
									),
									htmx.HxValidate(true),
									htmx.ID("operator-skgs"),
									htmx.Name("operator_skgs_id"),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Name"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:        "name",
										Placeholder: "Jarvis, Skynet, etc.",
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												tailwind.TextNeutral500: true,
											},
										},
										htmx.Text("The name must be from 3 to 255 characters. At least 3 characters must be non-whitespace."),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												tailwind.TextNeutral500: true,
											},
										},
										htmx.Text("A brief description of the acount to provide context."),
									),
								),
								forms.TextareaBordered(
									forms.TextareaProps{
										Name:        "description",
										Placeholder: "A super cool tool that does amazing things ...",
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("The description must be from 3 to 1024 characters."),
									),
								),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Limits"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Enable JetStream"),
									),
								),
								forms.Toggle(
									forms.ToggleProps{
										Name:    "jetstream_enable",
										Checked: true,
										Value:   conv.String(true),
									},
									htmx.ID("jetstream_enable"),
								),
							),
							forms.FormControl(
								forms.FormControlProps{
									ClassNames: htmx.ClassNames{
										"py-4": true,
									},
								},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{},
										},
										htmx.Text("Max Disk Storage"),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex": true,
									},
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:  "jetstream_max_disk_storage",
											Value: "2.5",
										},
									),
									forms.SelectBordered(
										forms.SelectProps{
											ClassNames: htmx.ClassNames{
												"w-full": false,
											},
										},
										htmx.Name("jetstream_max_disk_storage_unit"),
										forms.Option(
											forms.OptionProps{
												Value: "B",
											},
											htmx.Text("Bytes"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "KB",
											},
											htmx.Text("KiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "MB",
											},
											htmx.Text("MiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value:    "GB",
												Selected: true,
											},
											htmx.Text("GB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "TB",
											},
											htmx.Text("TiB"),
										),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Streams"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:  "jetstream_max_streams",
										Value: "10",
									},
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Consumers"),
									),
								),
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:  "jetstream_max_consumers",
										Value: "10",
									},
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"w-full": true,
											},
										},
										htmx.Text("Max Stream Size"),
									),
								),
								htmx.Div(
									htmx.ClassNames{
										"flex": true,
									},
									forms.TextInputBordered(
										forms.TextInputProps{
											Name:  "jetstream_max_stream_size",
											Value: "2.6",
										},
									),
									forms.SelectBordered(
										forms.SelectProps{
											ClassNames: htmx.ClassNames{
												"w-full": false,
											},
										},
										htmx.Name("jetstream_max_stream_size_unit"),
										forms.Option(
											forms.OptionProps{
												Value: "bytes",
											},
											htmx.Text("Bytes"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "kilobit",
											},
											htmx.Text("KiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "megabit",
											},
											htmx.Text("MiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value:    "gigabit",
												Selected: true,
											},
											htmx.Text("GiB"),
										),
										forms.Option(
											forms.OptionProps{
												Value: "terabit",
											},
											htmx.Text("TiB"),
										),
									),
								),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{},
										htmx.Text("Require Max Bytes"),
									),
								),
								forms.Toggle(
									forms.ToggleProps{
										Checked: true,
									},
									htmx.Name("jetstream_max_bytes_required"),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Action("."),
									htmx.HxPost("/accounts/create"),
									htmx.HxTargetError("#alerts"),
									htmx.Attribute("type", "submit"),
									htmx.Text("Create Account"),
								),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								tailwind.M2: true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Account Server"),
							),
							forms.FormControl(
								forms.FormControlProps{},
								forms.TextInputBordered(
									forms.TextInputProps{
										Name:        "account_server_url",
										Placeholder: "https://example.com:8080",
									},
								),
								forms.FormControlLabel(
									forms.FormControlLabelProps{},
									forms.FormControlLabelText(
										forms.FormControlLabelTextProps{
											ClassNames: htmx.ClassNames{
												"text-neutral-500": true,
											},
										},
										htmx.Text("A valid URL with a scheme of http or https. Certificates need be valid."),
									),
								),
							),
							cards.Actions(
								cards.ActionsProps{},
								buttons.Button(
									buttons.ButtonProps{},
									htmx.Attribute("type", "submit"),
									htmx.Text("Create Operator"),
								),
							),
						),
					),
					cards.CardBordered(
						cards.CardProps{
							ClassNames: htmx.ClassNames{
								"my-2": true,
								"mx-2": true,
							},
						},
						cards.Body(
							cards.BodyProps{},
							cards.Title(
								cards.TitleProps{},
								htmx.Text("Tags - Optional"),
							),
							htmx.Div(
								alpine.XData(`{
						tags: [],
						addTag(tag) {
						  this.tags.push({name: '', value: ''});
						},
						removeTag(index) {
						  this.tags.splice(index, 1);
						}
					  }`),
								htmx.Template(
									alpine.XFor("(tag, index) in tags"),
									htmx.Attribute(":key", "index"),
									htmx.Div(
										htmx.ClassNames{
											tailwind.Flex:    true,
											tailwind.SpaceX4: true,
										},
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{},
											},
											forms.TextInputBordered(
												forms.TextInputProps{},
												alpine.XModel("tag.name"),
												alpine.XBind("name", "`tags.${index}.name`"),
											),
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{
														ClassNames: htmx.ClassNames{
															"text-neutral-500": true,
														},
													},
													htmx.Text("Key is a unique identifier. At least 3 characters must be non-whitespace."),
												),
											),
										),
										forms.FormControl(
											forms.FormControlProps{
												ClassNames: htmx.ClassNames{},
											},
											forms.TextInputBordered(
												forms.TextInputProps{},
												alpine.XModel("tag.value"),
												alpine.XBind("name", "`tags.${index}.value`"),
											),
											forms.FormControlLabel(
												forms.FormControlLabelProps{},
												forms.FormControlLabelText(
													forms.FormControlLabelTextProps{
														ClassNames: htmx.ClassNames{
															"text-neutral-500": true,
														},
													},
													htmx.Text("Value is a unique value for the key."),
												),
											),
										),
										buttons.Button(
											buttons.ButtonProps{
												Type: "button",
											},
											alpine.XOn("click", "removeTag(index)"),
											icons.TrashOutline(
												icons.IconProps{},
											),
										),
									),
								),
								cards.Actions(
									cards.ActionsProps{},
									buttons.Button(
										buttons.ButtonProps{
											Type: "button",
										},
										alpine.XOn("click", "addTag()"),
										htmx.Text("Add Tag"),
									),
								),
							),
						),
					),
				)
			},
		),
	)
}
