#include "driver/gpio.h"
#include "esp_err.h"
#include "freertos/FreeRTOS.h"
#include "freertos/task.h"

namespace tbots {

constexpr gpio_num_t LED_PIN = GPIO_NUM_1;

class thunderloop {

public:
  thunderloop() {
    const gpio_config_t config = {
        .pin_bit_mask = 1ULL << LED_PIN,
        .mode = GPIO_MODE_OUTPUT,
        .pull_up_en = GPIO_PULLUP_DISABLE,
        .pull_down_en = GPIO_PULLDOWN_DISABLE,
        .intr_type = GPIO_INTR_DISABLE,
    };

    ESP_ERROR_CHECK(gpio_config(&config));
  }

  void run() {
    while (true) {
      vTaskDelay(1000 / portTICK_PERIOD_MS);
      ESP_ERROR_CHECK(gpio_set_level(LED_PIN, 1));
      vTaskDelay(1000 / portTICK_PERIOD_MS);
      ESP_ERROR_CHECK(gpio_set_level(LED_PIN, 0));
    }
  }
};

} // namespace tbots
