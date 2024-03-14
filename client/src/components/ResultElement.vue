<div>
      <div class="accordion mt-4" role="tablist">
        <b-card no-body class="mb-1">
          <b-card-header header-tag="header" class="p-1" role="tab">
            <b-button block v-b-toggle.accordion-1 variant="info"
              >Results</b-button
            >
          </b-card-header>
          <b-collapse
            id="accordion-1"
            visible
            accordion="my-accordion"
            role="tabpanel"
          >
            <b-card-body>
              <b-table hover :items="searches" :fields="fields">
                <template #cell(actions)="row">
                  <b-dropdown
                    id="dropdown-1"
                    text="Dropdown Button"
                    class="m-md-2"
                  >
                    <b-dropdown-item @click="viewSearch(row.item)"
                      >View</b-dropdown-item
                    >
                    <b-dropdown-item @click="enhanceSearch(row.item)">
                      Enhance</b-dropdown-item
                    >
                    <b-dropdown-item @click="saveSearch(row.item)"
                      >Save</b-dropdown-item
                    >
                    <b-dropdown-item @click="deleteSearch(row.item)"
                      >Delete</b-dropdown-item
                    >
                    <b-dropdown-divider></b-dropdown-divider>
                  </b-dropdown>
                </template>
              </b-table>
            </b-card-body>
          </b-collapse>
        </b-card>
      </div>
    </div>