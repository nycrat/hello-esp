#include "thunderloop.hpp"

extern "C" void app_main() {
  tbots::thunderloop thunderloop;
  thunderloop.run();
}
