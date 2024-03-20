#include "simple_factory.h"

std::string ProductA::Operation() {
    return "Product A";
}

std::string ProductB::Operation() {
    return "Product B";
}

Product* GetProduct(const std::string& productType) {
    if (productType == "A") {
        return new ProductA();
    } else if (productType == "B") {
        return new ProductB();
    }
    return nullptr;
}