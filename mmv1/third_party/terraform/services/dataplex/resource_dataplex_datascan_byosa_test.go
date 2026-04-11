package dataplex_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataplexDatascan_byosa(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexDatascanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascan_byosa(context),
			},
			{
				ResourceName:      "google_dataplex_datascan.byosa",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"data_scan_id", "location", "terraform_labels", "labels"},
			},
		},
	})
}

func testAccDataplexDatascan_byosa(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_account" "sa" {
  account_id   = "tf-test-datascan-sa%{random_suffix}"
  display_name = "DataScan Service Account"
  project      = "%{project_name}"
}

resource "google_dataplex_datascan" "byosa" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan-byosa%{random_suffix}"

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  execution_identity {
    service_account {
      email = google_service_account.sa.email
    }
  }

  data_profile_spec {}

  project = "%{project_name}"
}
`, context)
}

func TestAccDataplexDatascan_userCredential(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexDatascanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascan_userCredential(context),
			},
			{
				ResourceName:      "google_dataplex_datascan.user_cred",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"data_scan_id", "location", "terraform_labels", "labels"},
			},
		},
	})
}

func testAccDataplexDatascan_userCredential(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_datascan" "user_cred" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan-usercred%{random_suffix}"

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  execution_identity {
    user_credential {}
  }

  data_profile_spec {}

  project = "%{project_name}"
}
`, context)
}

func TestAccDataplexDatascan_dataplexServiceAgent(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexDatascanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascan_dataplexServiceAgent(context),
			},
			{
				ResourceName:      "google_dataplex_datascan.service_agent",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"data_scan_id", "location", "terraform_labels", "labels"},
			},
		},
	})
}

func testAccDataplexDatascan_dataplexServiceAgent(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_datascan" "service_agent" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan-serviceagent%{random_suffix}"

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  execution_identity {
    dataplex_service_agent {}
  }

  data_profile_spec {}

  project = "%{project_name}"
}
`, context)
}
package dataplex_test

import (
	"testing"

	"github.com/hashicorp/terraform-plugin-testing/helper/resource"

	"github.com/hashicorp/terraform-provider-google/google/acctest"
	"github.com/hashicorp/terraform-provider-google/google/envvar"
)

func TestAccDataplexDatascan_byosa(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexDatascanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascan_byosa(context),
			},
			{
				ResourceName:      "google_dataplex_datascan.byosa",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"data_scan_id", "location", "terraform_labels", "labels"},
			},
		},
	})
}

func testAccDataplexDatascan_byosa(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_service_account" "sa" {
  account_id   = "tf-test-datascan-sa%{random_suffix}"
  display_name = "DataScan Service Account"
  project      = "%{project_name}"
}

resource "google_dataplex_datascan" "byosa" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan-byosa%{random_suffix}"

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  execution_identity {
    service_account {
      email = google_service_account.sa.email
    }
  }

  data_profile_spec {}

  project = "%{project_name}"
}
`, context)
}

func TestAccDataplexDatascan_userCredential(t *testing.T) {
	t.Parallel()

	context := map[string]interface{}{
		"project_name":  envvar.GetTestProjectFromEnv(),
		"random_suffix": acctest.RandString(t, 10),
	}

	acctest.VcrTest(t, resource.TestCase{
		PreCheck:                 func() { acctest.AccTestPreCheck(t) },
		ProtoV5ProviderFactories: acctest.ProtoV5ProviderFactories(t),
		CheckDestroy:             testAccCheckDataplexDatascanDestroyProducer(t),
		Steps: []resource.TestStep{
			{
				Config: testAccDataplexDatascan_userCredential(context),
			},
			{
				ResourceName:      "google_dataplex_datascan.user_cred",
				ImportState:       true,
				ImportStateVerify: true,
				ImportStateVerifyIgnore: []string{"data_scan_id", "location", "terraform_labels", "labels"},
			},
		},
	})
}

func testAccDataplexDatascan_userCredential(context map[string]interface{}) string {
	return acctest.Nprintf(`
resource "google_dataplex_datascan" "user_cred" {
  location     = "us-central1"
  data_scan_id = "tf-test-datascan-usercred%{random_suffix}"

  data {
    resource = "//bigquery.googleapis.com/projects/bigquery-public-data/datasets/samples/tables/shakespeare"
  }

  execution_spec {
    trigger {
      on_demand {}
    }
  }

  execution_identity {
    user_credential {}
  }

  data_profile_spec {}

  project = "%{project_name}"
}
`, context)
}
