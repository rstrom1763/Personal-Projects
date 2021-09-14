import matplotlib.pyplot as plt


def calculate_interest(principal, rate, contribution):
    return (principal + contribution + (principal + contribution) * (rate / 100))


initial = 0
principal = initial
rate = 10
contribution = 6000
data = []
years = range(40)
for i in years:
    principal = calculate_interest(principal, rate, contribution)
    data.append(principal/1000000)

print(round((data[len(data)-1]), 3))

plt.plot(years, data, linewidth=3.0, color="red")

plt.axhline(y=0, linewidth=0.5, color="black")
plt.title('\$' + str(contribution) + ' annually, ' +
          str(rate) + ' percent annual interest')

plt.xlabel('Year')
plt.ylabel('Dollars (Millions)')
plt.show()
