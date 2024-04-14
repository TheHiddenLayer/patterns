#include <iostream>
#include <iterator>
#include <string>
#include <vector>

// Product is a concrete product
class Product
{
public:
    std::vector<std::string> parts_;
    // ListParts prints the contents of `parts_` to stdout
    void ListParts()
    {
        for (const auto& part : parts_)
        {
            std::cout << part << " ";
        }
        std::cout << "\n";
    }
};

// Builder is an interface that concrete builders ought to implement
class Builder
{
public:
    virtual ~Builder(){};

    // the various builder methods which each, individually their respective
    // parts of the final product. Note that this builder doens't have a Build()
    virtual void PartA() const = 0;
    virtual void PartB() const = 0;
    virtual void PartC() const = 0;
};

// ConcreteBuilder1 is a concrete builder
class ConcreteBuilder1 : public Builder
{
private:
    Product* product;

public:
    ConcreteBuilder1()
    {
        this->Reset();
    }

    ~ConcreteBuilder1()
    {
        delete this->product;
    }

    void Reset()
    {
        this->product = new Product();
    }

    // concrete build steps
    void PartA() const override
    {
        this->product->parts_.emplace_back("PartA");
    }
    void PartB() const override
    {
        this->product->parts_.emplace_back("PartB");
    }
    void PartC() const override
    {
        this->product->parts_.emplace_back("PartC");
    }

    // TODO user smart pointers to prevent memory leak
    Product* Build()
    {
        Product* product = this->product;
        this->Reset();
        return product;
    }
};

// Director has the responsibility of driving the construction steps,
// providing the client with a simple interface to create a product.
// NOTE: this implementation of the director requires the client to finally
// call Build() on the builder itself, rather than the director.
class Director
{
private:
    Builder* builder;

public:
    void SetBuilder(Builder* builder)
    {
        this->builder = builder;
    }

    // minimum viable product
    void BuildMVP()
    {
        this->builder->PartA();
    }

    // fully featured product
    void BuildFullProduct()
    {
        this->builder->PartA();
        this->builder->PartB();
        this->builder->PartC();
    }
};

int main()
{
    ConcreteBuilder1* builder = new ConcreteBuilder1();
    Director* d = new Director();
    d->SetBuilder(builder);
    d->BuildMVP();

    auto product = builder->Build();
    product->ListParts();

    d->BuildFullProduct();
    product = builder->Build();
    product->ListParts();
}
